/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/dict.go                                         |
|                                                          |
| LastModified: Feb 27, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"strconv"
	"sync"
)

// Dict represent the key-value pairs.
type Dict interface {
	Set(key string, value interface{})
	Get(key string) (value interface{}, ok bool)
	GetInt(key string, defaultValue ...int) int
	GetUInt(key string, defaultValue ...uint) uint
	GetInt64(key string, defaultValue ...int64) int64
	GetUInt64(key string, defaultValue ...uint64) uint64
	GetFloat(key string, defaultValue ...float64) float64
	GetBool(key string, defaultValue ...bool) bool
	GetString(key string, defaultValue ...string) string
	GetInterface(key string, defaultValue ...interface{}) interface{}
	Del(key string)
	Range(f func(key string, value interface{}) bool)
	Empty() bool
	CopyTo(dict Dict)
	ToMap() map[string]interface{}
}

func getInt(d Dict, key string, defaultValue ...int) int {
	if value, ok := d.Get(key); ok {
		switch value := value.(type) {
		case int:
			return value
		case uint:
			return int(value)
		case int64:
			return int(value)
		case uint64:
			return int(value)
		case int32:
			return int(value)
		case uint32:
			return int(value)
		case int16:
			return int(value)
		case uint16:
			return int(value)
		case int8:
			return int(value)
		case uint8:
			return int(value)
		case float64:
			return int(value)
		case float32:
			return int(value)
		case bool:
			if value {
				return 1
			}
			return 0
		case string:
			if i, err := strconv.ParseInt(value, 0, 0); err == nil {
				return int(i)
			}
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

func getUInt(d Dict, key string, defaultValue ...uint) uint {
	if value, ok := d.Get(key); ok {
		switch value := value.(type) {
		case uint:
			return value
		case int:
			return uint(value)
		case uint64:
			return uint(value)
		case int64:
			return uint(value)
		case uint32:
			return uint(value)
		case int32:
			return uint(value)
		case uint16:
			return uint(value)
		case int16:
			return uint(value)
		case uint8:
			return uint(value)
		case int8:
			return uint(value)
		case float64:
			return uint(value)
		case float32:
			return uint(value)
		case bool:
			if value {
				return 1
			}
			return 0
		case string:
			if i, err := strconv.ParseUint(value, 0, 0); err == nil {
				return uint(i)
			}
		}
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

func getInt64(d Dict, key string, defaultValue ...int64) int64 {
	if value, ok := d.Get(key); ok {
		switch value := value.(type) {
		case int64:
			return value
		case int:
			return int64(value)
		case uint:
			return int64(value)
		case uint64:
			return int64(value)
		case int32:
			return int64(value)
		case uint32:
			return int64(value)
		case int16:
			return int64(value)
		case uint16:
			return int64(value)
		case int8:
			return int64(value)
		case uint8:
			return int64(value)
		case float64:
			return int64(value)
		case float32:
			return int64(value)
		case bool:
			if value {
				return 1
			}
			return 0
		case string:
			if i, err := strconv.ParseInt(value, 0, 0); err == nil {
				return i
			}
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

func getUInt64(d Dict, key string, defaultValue ...uint64) uint64 {
	if value, ok := d.Get(key); ok {
		switch value := value.(type) {
		case uint64:
			return value
		case uint:
			return uint64(value)
		case int:
			return uint64(value)
		case int64:
			return uint64(value)
		case uint32:
			return uint64(value)
		case int32:
			return uint64(value)
		case uint16:
			return uint64(value)
		case int16:
			return uint64(value)
		case uint8:
			return uint64(value)
		case int8:
			return uint64(value)
		case float64:
			return uint64(value)
		case float32:
			return uint64(value)
		case bool:
			if value {
				return 1
			}
			return 0
		case string:
			if i, err := strconv.ParseUint(value, 0, 0); err == nil {
				return i
			}
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

func getFloat(d Dict, key string, defaultValue ...float64) float64 {
	if value, ok := d.Get(key); ok {
		switch value := value.(type) {
		case float64:
			return value
		case float32:
			return float64(value)
		case int:
			return float64(value)
		case uint:
			return float64(value)
		case int64:
			return float64(value)
		case uint64:
			return float64(value)
		case int32:
			return float64(value)
		case uint32:
			return float64(value)
		case int16:
			return float64(value)
		case uint16:
			return float64(value)
		case int8:
			return float64(value)
		case uint8:
			return float64(value)
		case bool:
			if value {
				return 1
			}
			return 0
		case string:
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				return f
			}
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

func getBool(d Dict, key string, defaultValue ...bool) bool {
	if value, ok := d.Get(key); ok {
		switch value := value.(type) {
		case bool:
			return value
		case float64:
			return value != 0
		case float32:
			return value != 0
		case int:
			return value != 0
		case uint:
			return value != 0
		case int64:
			return value != 0
		case uint64:
			return value != 0
		case int32:
			return value != 0
		case uint32:
			return value != 0
		case int16:
			return value != 0
		case uint16:
			return value != 0
		case int8:
			return value != 0
		case uint8:
			return value != 0
		case string:
			if b, err := strconv.ParseBool(value); err == nil {
				return b
			}
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return false
}

func getString(d Dict, key string, defaultValue ...string) string {
	if value, ok := d.Get(key); ok {
		switch value := value.(type) {
		case string:
			return value
		case bool:
			return strconv.FormatBool(value)
		case float64:
			return strconv.FormatFloat(value, 'g', -1, 64)
		case float32:
			return strconv.FormatFloat(float64(value), 'g', -1, 32)
		case int:
			return strconv.FormatInt(int64(value), 10)
		case uint:
			return strconv.FormatUint(uint64(value), 10)
		case int64:
			return strconv.FormatInt(value, 10)
		case uint64:
			return strconv.FormatUint(value, 10)
		case int32:
			return strconv.FormatInt(int64(value), 10)
		case uint32:
			return strconv.FormatUint(uint64(value), 10)
		case int16:
			return strconv.FormatInt(int64(value), 10)
		case uint16:
			return strconv.FormatUint(uint64(value), 10)
		case int8:
			return strconv.FormatInt(int64(value), 10)
		case uint8:
			return strconv.FormatUint(uint64(value), 10)
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

func getInterface(d Dict, key string, defaultValue ...interface{}) interface{} {
	if value, ok := d.Get(key); ok {
		return value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return nil
}

type dict map[string]interface{}

func (d dict) Set(key string, value interface{}) {
	d[key] = value
}

func (d dict) Get(key string) (value interface{}, ok bool) {
	value, ok = d[key]
	return
}

func (d dict) GetInt(key string, defaultValue ...int) int {
	return getInt(d, key, defaultValue...)
}

func (d dict) GetUInt(key string, defaultValue ...uint) uint {
	return getUInt(d, key, defaultValue...)
}

func (d dict) GetInt64(key string, defaultValue ...int64) int64 {
	return getInt64(d, key, defaultValue...)
}

func (d dict) GetUInt64(key string, defaultValue ...uint64) uint64 {
	return getUInt64(d, key, defaultValue...)
}

func (d dict) GetFloat(key string, defaultValue ...float64) float64 {
	return getFloat(d, key, defaultValue...)
}

func (d dict) GetBool(key string, defaultValue ...bool) bool {
	return getBool(d, key, defaultValue...)
}

func (d dict) GetString(key string, defaultValue ...string) string {
	return getString(d, key, defaultValue...)
}

func (d dict) GetInterface(key string, defaultValue ...interface{}) interface{} {
	return getInterface(d, key, defaultValue...)
}

func (d dict) Del(key string) {
	delete(d, key)
}

func (d dict) Range(f func(key string, value interface{}) bool) {
	for k, v := range d {
		if !f(k, v) {
			return
		}
	}
}

func (d dict) Empty() bool {
	return len(d) == 0
}

func (d dict) CopyTo(dict Dict) {
	for k, v := range d {
		dict.Set(k, v)
	}
}

func (d dict) ToMap() map[string]interface{} {
	return d
}

// NewDict returns a thread-unsafe Dict.
func NewDict(m map[string]interface{}) Dict {
	if m == nil {
		return dict(make(map[string]interface{}))
	}
	return dict(m)
}

type safeDict sync.Map

func (d *safeDict) unwarp() *sync.Map {
	return (*sync.Map)(d)
}

func (d *safeDict) Set(name string, value interface{}) {
	d.unwarp().Store(name, value)
}

func (d *safeDict) Get(name string) (value interface{}, ok bool) {
	return d.unwarp().Load(name)
}

func (d *safeDict) GetInt(key string, defaultValue ...int) int {
	return getInt(d, key, defaultValue...)
}

func (d *safeDict) GetUInt(key string, defaultValue ...uint) uint {
	return getUInt(d, key, defaultValue...)
}

func (d *safeDict) GetInt64(key string, defaultValue ...int64) int64 {
	return getInt64(d, key, defaultValue...)
}

func (d *safeDict) GetUInt64(key string, defaultValue ...uint64) uint64 {
	return getUInt64(d, key, defaultValue...)
}

func (d *safeDict) GetFloat(key string, defaultValue ...float64) float64 {
	return getFloat(d, key, defaultValue...)
}

func (d *safeDict) GetBool(key string, defaultValue ...bool) bool {
	return getBool(d, key, defaultValue...)
}

func (d *safeDict) GetString(key string, defaultValue ...string) string {
	return getString(d, key, defaultValue...)
}

func (d *safeDict) GetInterface(key string, defaultValue ...interface{}) interface{} {
	return getInterface(d, key, defaultValue...)
}

func (d *safeDict) Del(name string) {
	d.unwarp().Delete(name)
}

func (d *safeDict) Range(f func(name string, value interface{}) bool) {
	d.unwarp().Range(func(key, value interface{}) bool {
		return f(key.(string), value)
	})
}

func (d *safeDict) Empty() bool {
	empty := true
	d.unwarp().Range(func(key, value interface{}) bool {
		empty = false
		return false
	})
	return empty
}

func (d *safeDict) CopyTo(dict Dict) {
	d.unwarp().Range(func(key, value interface{}) bool {
		dict.Set(key.(string), value)
		return true
	})
}

func (d *safeDict) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	d.unwarp().Range(func(key, value interface{}) bool {
		m[key.(string)] = value
		return true
	})
	return m
}

// NewSafeDict returns a thread-safe Dict.
func NewSafeDict() Dict {
	return &safeDict{}
}
