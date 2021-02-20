/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/proxy.go                                        |
|                                                          |
| LastModified: Feb 20, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"reflect"
)

// InvocationHandler for the proxy instance.
type InvocationHandler func(proxy interface{}, method reflect.StructField, name string, args []interface{}) (results []interface{}, err error)

// ProxyBuilder .
type ProxyBuilder interface {
	Build(proxy interface{}, handler InvocationHandler)
}

// Proxy is a global ProxyBuilder.
var Proxy ProxyBuilder = proxyBuilder{}

type proxyBuilder struct{}

func (b proxyBuilder) Build(proxy interface{}, handler InvocationHandler) {
	b.build(proxy, handler, "", reflect.ValueOf(proxy).Elem())
}

func (b proxyBuilder) build(proxy interface{}, handler InvocationHandler, namespace string, p reflect.Value) {
	t := p.Type()
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		if p.IsNil() {
			setAccessible(p).Set(reflect.New(t))
		}
		p = p.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}
	n := p.NumField()
	for i := 0; i < n; i++ {
		f := p.Field(i)
		ft := f.Type()
		sf := t.Field(i)
		name := namespace
		if !sf.Anonymous {
			name = sf.Name
			if namespace != "" {
				name = namespace + "." + sf.Name
			}
		}
		switch ft.Kind() {
		case reflect.Struct, reflect.Ptr:
			b.build(proxy, handler, name, f)
		case reflect.Func:
			setAccessible(f).Set(b.method(proxy, handler, name, ft, sf))
		}
	}
}

func (b proxyBuilder) in(ft reflect.Type, in []reflect.Value) (args []interface{}) {
	n := len(in)
	m := n - 1
	if ft.IsVariadic() {
		n--
		n += in[n].Len()
	}
	args = make([]interface{}, n)
	if ft.IsVariadic() {
		for i := 0; i < m; i++ {
			args[i] = in[i].Interface()
		}
		v := in[m]
		for i := 0; i < v.Len(); i++ {
			args[m+i] = v.Index(i).Interface()
		}
	} else {
		for i := 0; i < n; i++ {
			args[i] = in[i].Interface()
		}
	}
	return
}

func (b proxyBuilder) out(ft reflect.Type, results []interface{}, err error) (out []reflect.Value) {
	n := ft.NumOut()
	out = make([]reflect.Value, n)
	if ft.Out(n-1) == errorType {
		n--
		out[n] = reflect.ValueOf(&err).Elem()
	}
	m := len(results)
	if m > n {
		m = n
	}
	for i := 0; i < m; i++ {
		out[i] = reflect.ValueOf(results[i])
	}
	for i := m; i < n; i++ {
		out[i] = reflect.Zero(ft.Out(i))
	}
	return
}

func (b proxyBuilder) method(proxy interface{}, h InvocationHandler, name string, ft reflect.Type, sf reflect.StructField) reflect.Value {
	return reflect.MakeFunc(ft, func(in []reflect.Value) (out []reflect.Value) {
		args := b.in(ft, in)
		results, err := h(proxy, sf, name, args)
		return b.out(ft, results, err)
	})
}
