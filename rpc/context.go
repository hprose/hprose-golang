/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * rpc/context.go                                         *
 *                                                        *
 * hprose context for Go.                                 *
 *                                                        *
 * LastModified: Dec 1, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import "reflect"

// Context is the hprose context
type Context interface {
	UserData() map[string]interface{}
	GetInt(key string, defaultValue ...int) int
	GetUInt(key string, defaultValue ...uint) uint
	GetInt64(key string, defaultValue ...int64) int64
	GetUInt64(key string, defaultValue ...uint64) uint64
	GetFloat(key string, defaultValue ...float64) float64
	GetBool(key string, defaultValue ...bool) bool
	GetString(key string, defaultValue ...string) string
	GetInterface(key string, defaultValue ...interface{}) interface{}
	Get(key string) interface{}
	SetInt(key string, value int)
	SetUInt(key string, value uint)
	SetInt64(key string, value int64)
	SetUInt64(key string, value uint64)
	SetFloat(key string, value float64)
	SetBool(key string, value bool)
	SetString(key string, value string)
	SetInterface(key string, value interface{})
	Set(key string, value interface{})
}

// BaseContext is the base context
type BaseContext struct {
	userData map[string]interface{}
}

// InitBaseContext initializes BaseContext
func (context *BaseContext) InitBaseContext() {
	context.userData = make(map[string]interface{})
}

// UserData return the user data
func (context *BaseContext) UserData() map[string]interface{} {
	return context.userData
}

// GetInt from hprose context
func (context *BaseContext) GetInt(
	key string, defaultValue ...int) int {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(int); ok {
			return value
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// GetUInt from hprose context
func (context *BaseContext) GetUInt(
	key string, defaultValue ...uint) uint {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(uint); ok {
			return value
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// GetInt64 from hprose context
func (context *BaseContext) GetInt64(
	key string, defaultValue ...int64) int64 {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(int64); ok {
			return value
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// GetUInt64 from hprose context
func (context *BaseContext) GetUInt64(
	key string, defaultValue ...uint64) uint64 {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(uint64); ok {
			return value
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// GetFloat from hprose context
func (context *BaseContext) GetFloat(
	key string, defaultValue ...float64) float64 {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(float64); ok {
			return value
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// GetBool from hprose context
func (context *BaseContext) GetBool(
	key string, defaultValue ...bool) bool {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(bool); ok {
			return value
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return false
}

// GetString from hprose context
func (context *BaseContext) GetString(
	key string, defaultValue ...string) string {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(string); ok {
			return value
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

// GetInterface from hprose context
func (context *BaseContext) GetInterface(
	key string, defaultValue ...interface{}) interface{} {
	if value, ok := context.userData[key]; ok {
		return value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return nil
}

// Get value from hprose context
func (context *BaseContext) Get(key string) interface{} {
	if value, ok := context.userData[key]; ok {
		return value
	}
	return nil
}

func (context *BaseContext) checkType(val1 interface{}, val2 interface{}) bool {
	refTyp1 := reflect.TypeOf(val1)
	refTyp2 := reflect.TypeOf(val2)
	if refTyp1.Kind() == refTyp2.Kind() {
		return true
	}
	return false
}

// SetInt to hprose context
func (context *BaseContext) SetInt(key string, value int) {
	if oldVal, ok := context.userData[key]; !ok {
		context.userData[key] = value
		return
	} else if !context.checkType(oldVal, value) {
		panic("context set value's type is not right")
	}
	context.userData[key] = value
	return
}

// SetUInt to hprose context
func (context *BaseContext) SetUInt(key string, value uint) {
	if oldVal, ok := context.userData[key]; !ok {
		context.userData[key] = value
		return
	} else if !context.checkType(oldVal, value) {
		panic("context set value's type is not right")
	}
	context.userData[key] = value
	return
}

// SetInt64 to hprose context
func (context *BaseContext) SetInt64(key string, value int64) {
	if oldVal, ok := context.userData[key]; !ok {
		context.userData[key] = value
		return
	} else if !context.checkType(oldVal, value) {
		panic("context set value's type is not right")
	}
	context.userData[key] = value
}

// SetUInt64 to hprose context
func (context *BaseContext) SetUInt64(key string, value uint64) {
	if oldVal, ok := context.userData[key]; !ok {
		context.userData[key] = value
		return
	} else if !context.checkType(oldVal, value) {
		panic("context set value's type is not right")
	}
	context.userData[key] = value
	return
}

// SetFloat to hprose context
func (context *BaseContext) SetFloat(key string, value float64) {
	if oldVal, ok := context.userData[key]; !ok {
		context.userData[key] = value
		return
	} else if !context.checkType(oldVal, value) {
		panic("context set value's type is not right")
	}
	context.userData[key] = value
	return
}

// SetBool to hprose context
func (context *BaseContext) SetBool(key string, value bool) {
	if oldVal, ok := context.userData[key]; !ok {
		context.userData[key] = value
		return
	} else if !context.checkType(oldVal, value) {
		panic("context set value's type is not right")
	}
	context.userData[key] = value
	return
}

// SetString to hprose context
func (context *BaseContext) SetString(key string, value string) {
	if oldVal, ok := context.userData[key]; !ok {
		context.userData[key] = value
		return
	} else if !context.checkType(oldVal, value) {
		panic("context set value's type is not right")
	}
	context.userData[key] = value
	return
}

// SetInterface to hprose context
func (context *BaseContext) SetInterface(key string, value interface{}) {
	if oldVal, ok := context.userData[key]; !ok {
		context.userData[key] = value
		return
	} else if !context.checkType(oldVal, value) {
		panic("context set value's type is not right")
	}
	context.userData[key] = value
	return
}

// Set is an alias of SetInterface
func (context *BaseContext) Set(key string, value interface{}) {
	if oldVal, ok := context.userData[key]; !ok {
		context.userData[key] = value
		return
	} else if !context.checkType(oldVal, value) {
		panic("context set value's type is not right")
	}
	context.userData[key] = value
	return
}
