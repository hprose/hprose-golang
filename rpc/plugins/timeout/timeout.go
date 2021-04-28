/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/timeout/timeout.go                           |
|                                                          |
| LastModified: Feb 21, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package timeout

import (
	"context"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// ExecuteTimeout plugin for hprose.
type ExecuteTimeout struct {
	Timeout time.Duration
}

type returnValue struct {
	result []interface{}
	err    error
}

// Handler for ExecuteTimeout.
func (et *ExecuteTimeout) Handler(ctx context.Context, name string, args []interface{}, next core.NextInvokeHandler) (result []interface{}, err error) {
	timeout := et.Timeout
	serviceContext := core.GetServiceContext(ctx)
	if t, ok := serviceContext.Method.Options().Get("timeout"); ok {
		switch t := t.(type) {
		case time.Duration:
			timeout = t
		case int:
			timeout = time.Duration(t)
		case uint:
			timeout = time.Duration(t)
		case int64:
			timeout = time.Duration(t)
		case uint64:
			timeout = time.Duration(t)
		}
	}
	if timeout <= 0 {
		return next(ctx, name, args)
	}
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()
	c := make(chan returnValue, 1)
	go func() {
		result, err := next(ctx, name, args)
		c <- returnValue{result, err}
	}()
	select {
	case <-ctx.Done():
		return nil, core.ErrTimeout
	case r := <-c:
		return r.result, r.err
	}
}

// New returns an ExecuteTimeout instance.
func New(timeout ...time.Duration) *ExecuteTimeout {
	if len(timeout) > 0 {
		return &ExecuteTimeout{timeout[0]}
	}
	return &ExecuteTimeout{time.Second * 30}
}
