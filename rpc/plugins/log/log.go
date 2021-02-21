/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/log/log.go                                   |
|                                                          |
| LastModified: Feb 21, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package log

import (
	"context"
	"fmt"
	"unsafe"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type Log struct {
	Println func(v interface{})
	Enabled bool
}

// NewLog returns a Log instance.
func NewLog(f ...func(v interface{})) *Log {
	p := func(v interface{}) { fmt.Println(v) }
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

// IOHandler for log.
func (log *Log) IOHandler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	enabled := log.Enabled
	if context, ok := core.FromContext(ctx); ok {
		enabled = context.Items().GetBool("log", enabled)
	}
	if !enabled {
		return next(ctx, request)
	}
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
			panic(e)
		}
		if err != nil {
			log.Println(err)
		} else {
			log.Println(unsafeString(response))
		}
	}()
	log.Println(unsafeString(request))
	return next(ctx, request)
}

var log = NewLog()

// IOHandler is the default io handler for log.
func IOHandler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	return log.IOHandler(ctx, request, next)
}
