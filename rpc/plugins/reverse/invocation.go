/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/reverse/invocation.go                        |
|                                                          |
| LastModified: May 19, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package reverse

import (
	"context"
	"reflect"
	"strings"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

var errorType = reflect.TypeOf((*error)(nil)).Elem()

type invocation struct {
	caller    *Caller
	id        string
	namespace string
}

func (i invocation) Invoke(proxy interface{}, method reflect.StructField, name string, args []interface{}) (results []interface{}, err error) {
	var rpcContext core.Context
	var ctx context.Context
	if len(args) > 0 {
		switch c := args[0].(type) {
		case core.Context:
			rpcContext = c
			args = args[1:]
		case context.Context:
			ctx = c
			rpcContext, _ = core.FromContext(ctx)
			args = args[1:]
		}
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if rpcContext == nil {
		rpcContext = core.NewContext()
	}
	if _, ok := core.FromContext(ctx); !ok {
		ctx = core.WithContext(ctx, rpcContext)
	}
	tagParser := core.ParseTag(nil, method.Tag)
	if tagParser.Name != "" {
		name = tagParser.Name
	}
	name = strings.Replace(name, ".", "_", -1) //nolint:gocritic
	if i.namespace != "" {
		name = i.namespace + "_" + name
	}
	t := method.Type
	n := t.NumOut()
	returnType := make([]reflect.Type, n)
	for i := 0; i < n; i++ {
		returnType[i] = t.Out(i)
	}
	if n > 0 && returnType[n-1] == errorType {
		returnType = returnType[:n-1]
	}
	return i.caller.InvokeContext(ctx, i.id, name, args, returnType...)
}
