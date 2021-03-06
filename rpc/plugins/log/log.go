/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/log/log.go                                   |
|                                                          |
| LastModified: Mar 6, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package log

import (
	"context"
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// Log plugin for hprose.
type Log struct {
	Println func(v ...interface{})
	Enabled bool
}

// New returns a Log instance.
func New(f ...func(v ...interface{})) *Log {
	p := func(v ...interface{}) { fmt.Println(v...) }
	if len(f) > 0 {
		p = f[0]
	}
	return &Log{
		Println: p,
		Enabled: true,
	}
}

func unsafeString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func (log *Log) isEnabled(ctx context.Context) (enabled bool) {
	enabled = log.Enabled
	if context, ok := core.FromContext(ctx); ok {
		enabled = context.Items().GetBool("log", enabled)
	}
	return
}

// IOHandler for log.
func (log *Log) IOHandler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	if !log.isEnabled(ctx) {
		return next(ctx, request)
	}
	defer func() {
		if e := recover(); e != nil {
			err = core.NewPanicError(e)
		}
		if err != nil {
			log.Println("error:", err)
		} else {
			log.Println("response:", unsafeString(response))
		}
	}()
	log.Println("request:", unsafeString(request))
	return next(ctx, request)
}

// InvokeHandler for log.
func (log *Log) InvokeHandler(ctx context.Context, name string, args []interface{}, next core.NextInvokeHandler) (result []interface{}, err error) {
	if !log.isEnabled(ctx) {
		return next(ctx, name, args)
	}
	defer func() {
		if e := recover(); e != nil {
			err = core.NewPanicError(e)
		}
		if err != nil {
			log.Println("error:", err)
		} else if data, e := json.Marshal(result); e == nil {
			log.Println("result:", unsafeString(data))
		} else {
			log.Println("result:", result)
		}
	}()
	log.Println("name:", name)
	if data, e := json.Marshal(args); e == nil {
		log.Println("args:", unsafeString(data))
	} else {
		log.Println("args:", args)
	}
	return next(ctx, name, args)
}

// Plugin is the default log plugin.
var Plugin = New()

// IOHandler is the default io handler for log.
func IOHandler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	return Plugin.IOHandler(ctx, request, next)
}

// InvokeHandler is the default io handler for log.
func InvokeHandler(ctx context.Context, name string, args []interface{}, next core.NextInvokeHandler) (result []interface{}, err error) {
	return Plugin.InvokeHandler(ctx, name, args, next)
}
