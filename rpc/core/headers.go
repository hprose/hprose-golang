/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/headers.go                                      |
|                                                          |
| LastModified: Jan 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import "sync"

// Headers represent the key-value pairs.
type Headers interface {
	Set(name string, value interface{})
	Get(name string) (value interface{}, ok bool)
	Del(name string)
	Range(func(name string, value interface{}) bool)
	Clone() Headers
	CopyTo(headers Headers)
}

type headers map[string]interface{}

func (h headers) Set(name string, value interface{}) {
	h[name] = value
}

func (h headers) Get(name string) (value interface{}, ok bool) {
	value, ok = h[name]
	return
}

func (h headers) Del(name string) {
	delete(h, name)
}

func (h headers) Range(f func(name string, value interface{}) bool) {
	for k, v := range h {
		if !f(k, v) {
			return
		}
	}
}

func (h headers) Clone() Headers {
	clone := headers(make(map[string]interface{}))
	for k, v := range h {
		clone[k] = v
	}
	return clone
}

func (h headers) CopyTo(headers Headers) {
	for k, v := range h {
		headers.Set(k, v)
	}
}

// NewHeaders returns a thread-unsafe Headers.
func NewHeaders() Headers {
	return headers(make(map[string]interface{}))
}

type safeHeaders struct {
	m *sync.Map
}

func (h safeHeaders) Set(name string, value interface{}) {
	h.m.Store(name, value)
}

func (h safeHeaders) Get(name string) (value interface{}, ok bool) {
	return h.m.Load(name)
}

func (h safeHeaders) Del(name string) {
	h.m.Delete(name)
}

func (h safeHeaders) Range(f func(name string, value interface{}) bool) {
	h.m.Range(func(key, value interface{}) bool {
		return f(key.(string), value)
	})
}

func (h safeHeaders) Clone() Headers {
	clone := safeHeaders{&sync.Map{}}
	h.m.Range(func(key, value interface{}) bool {
		clone.m.Store(key, value)
		return true
	})
	return clone
}

func (h safeHeaders) CopyTo(headers Headers) {
	h.m.Range(func(key, value interface{}) bool {
		headers.Set(key.(string), value)
		return true
	})
}

// NewSafeHeaders returns a thread-safe Headers.
func NewSafeHeaders() Headers {
	return safeHeaders{&sync.Map{}}
}
