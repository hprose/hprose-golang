/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/tag_parser.go                                   |
|                                                          |
| LastModified: Feb 20, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"reflect"
	"strconv"
	"time"
)

type tagParser struct {
	Name    string
	Context *ClientContext
	Tag     reflect.StructTag
}

func (tp *tagParser) parseName() {
	tp.Name = tp.Tag.Get("name")
}

func (tp *tagParser) parseTimeout() {
	if s, ok := tp.Tag.Lookup("timeout"); ok {
		if timeout, err := strconv.Atoi(s); err == nil {
			tp.Context.Timeout = time.Millisecond * time.Duration(timeout)
		}
	}
}

func (tp *tagParser) parseMapName(tag string) (remain string, name string, c byte) {
	// Skip leading space.
	i := 0
	for i < len(tag) && tag[i] == ' ' {
		i++
	}
	tag = tag[i:]
	if tag == "" {
		return
	}

	// Scan to colon or comma.
	i = 0
	for i < len(tag) && tag[i] != ':' && tag[i] != ',' {
		i++
	}
	c = tag[i]
	name = tag[:i]
	remain = tag[i+1:]
	return
}

func (tp *tagParser) parseMapValue(tag string) (string, string) {
	// Scan to find value.
	i := 0
	c := byte(',')
	if i < len(tag) && tag[i] == '"' {
		i++
		c = '"'
	} else if i < len(tag) && tag[i] == '\'' {
		i++
		c = '\''
	}
	for i < len(tag) && tag[i] != c {
		i++
	}
	if i < len(tag) && tag[i+1] == ',' {
		i++
	}
	value := tag[:i]
	tag = tag[i:]
	return tag, value
}

func (tp *tagParser) parseMap(key string) map[string]interface{} {
	m := make(map[string]interface{})
	tag := tp.Tag.Get(key)
	for tag != "" {
		var name string
		var c byte
		tag, name, c = tp.parseMapName(tag)
		if tag == "" {
			break
		}
		if c == ',' {
			m[name] = true
			continue
		}
		var value string
		tag, value = tp.parseMapValue(tag)
		if (len(value) >= 2) && (value[0] == '"' && value[len(value)-1] == '"') ||
			(value[0] == '\'' && value[len(value)-1] == '\'') {
			m[name] = value[1 : len(value)-1]
			continue
		}
		if value == "nil" || value == "null" {
			m[name] = nil
			continue
		}
		if intValue, err := strconv.Atoi(value); err == nil {
			m[name] = intValue
			continue
		}
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			m[name] = floatValue
			continue
		}
		if boolValue, err := strconv.ParseBool(value); err == nil {
			m[name] = boolValue
			continue
		}
		m[name] = value
	}
	return m
}

func (tp *tagParser) parseHeader() {
	m := tp.parseMap("header")
	header := tp.Context.RequestHeaders()
	for key, value := range m {
		header.Set(key, value)
	}
}

func (tp *tagParser) parseContext() {
	m := tp.parseMap("context")
	items := tp.Context.Items()
	for key, value := range m {
		items.Set(key, value)
	}
}

func (tp *tagParser) Parse() {
	tp.parseName()
	tp.parseTimeout()
	tp.parseHeader()
	tp.parseContext()
}
