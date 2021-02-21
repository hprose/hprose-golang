/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/invocation.go                                   |
|                                                          |
| LastModified: Feb 22, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"context"
	"reflect"
	"strings"
)

type invocation struct {
	client    *Client
	namespace string
}

func (i invocation) Invoke(proxy interface{}, method reflect.StructField, name string, args []interface{}) (results []interface{}, err error) {
	var clientContext *ClientContext
	var ctx context.Context
	if len(args) > 0 {
		switch c := args[0].(type) {
		case *ClientContext:
			clientContext = c
			args = args[1:]
		case *rpcContext:
			clientContext = &ClientContext{Context: c}
			args = args[1:]
		case context.Context:
			ctx = c
			clientContext = GetClientContext(c)
			args = args[1:]
		}
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if clientContext == nil {
		clientContext = NewClientContext()
	}
	if GetClientContext(ctx) == nil {
		ctx = WithContext(ctx, clientContext)
	}
	tagParser := parseTag(clientContext, method.Tag)
	if tagParser.Name != "" {
		name = tagParser.Name
	}
	name = strings.Replace(name, ".", "_", -1) //nolint:gocritic
	if i.namespace != "" {
		name = i.namespace + "_" + name
	}
	t := method.Type
	n := t.NumOut()
	clientContext.ReturnType = make([]reflect.Type, n)
	for i := 0; i < n; i++ {
		clientContext.ReturnType = append(clientContext.ReturnType, t.Out(i))
	}
	if n > 0 && clientContext.ReturnType[n-1] == errorType {
		clientContext.ReturnType = clientContext.ReturnType[:n-1]
	}
	return i.client.InvokeContext(ctx, name, args)
}
