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

// SetInt to hprose context
func (context *BaseContext) SetInt(key string, value int) {
	context.userData[key] = value
}

// SetUInt to hprose context
func (context *BaseContext) SetUInt(key string, value uint) {
	context.userData[key] = value
}

// SetInt64 to hprose context
func (context *BaseContext) SetInt64(key string, value int64) {
	context.userData[key] = value
}

// SetUInt64 to hprose context
func (context *BaseContext) SetUInt64(key string, value uint64) {
	context.userData[key] = value
}

// SetFloat to hprose context
func (context *BaseContext) SetFloat(key string, value float64) {
	context.userData[key] = value
}

// SetBool to hprose context
func (context *BaseContext) SetBool(key string, value bool) {
	context.userData[key] = value
}

// SetString to hprose context
func (context *BaseContext) SetString(key string, value string) {
	context.userData[key] = value
}

// SetInterface to hprose context
func (context *BaseContext) SetInterface(key string, value interface{}) {
	context.userData[key] = value
}

// Set is an alias of SetInterface
func (context *BaseContext) Set(key string, value interface{}) {
	context.userData[key] = value
}
