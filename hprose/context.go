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
 * hprose/context.go                                      *
 *                                                        *
 * hprose tcp service for Go.                             *
 *                                                        *
 * LastModified: Oct 15, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

type Context interface {
	UserData() map[string]interface{}
	GetInt(key string) (value int, ok bool)
	GetUInt(key string) (value uint, ok bool)
	GetInt64(key string) (value int64, ok bool)
	GetUInt64(key string) (value uint64, ok bool)
	GetFloat(key string) (value float64, ok bool)
	GetBool(key string) (value bool, ok bool)
	GetString(key string) (value string, ok bool)
	GetInterface(key string) (value interface{}, ok bool)
	SetInt(key string, value int)
	SetUInt(key string, value uint)
	SetInt64(key string, value int64)
	SetUInt64(key string, value uint64)
	SetFloat(key string, value float64)
	SetBool(key string, value bool)
	SetString(key string, value string)
	SetInterface(key string, value interface{})
}

type BaseContext struct {
	userData map[string]interface{}
}

func NewBaseContext() *BaseContext {
	return &BaseContext{userData: make(map[string]interface{})}
}

func (context *BaseContext) UserData() map[string]interface{} {
	return context.userData
}

func (context *BaseContext) GetInt(key string) (value int, ok bool) {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(int); ok {
			return value, true
		}
	}
	return 0, false
}

func (context *BaseContext) GetUInt(key string) (value uint, ok bool) {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(uint); ok {
			return value, true
		}
	}
	return 0, false
}

func (context *BaseContext) GetInt64(key string) (value int64, ok bool) {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(int64); ok {
			return value, true
		}
	}
	return 0, false
}

func (context *BaseContext) GetUInt64(key string) (value uint64, ok bool) {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(uint64); ok {
			return value, true
		}
	}
	return 0, false
}

func (context *BaseContext) GetFloat(key string) (value float64, ok bool) {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(float64); ok {
			return value, true
		}
	}
	return 0, false
}

func (context *BaseContext) GetBool(key string) (value bool, ok bool) {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(bool); ok {
			return value, true
		}
	}
	return false, false
}

func (context *BaseContext) GetString(key string) (value string, ok bool) {
	if value, ok := context.userData[key]; ok {
		if value, ok := value.(string); ok {
			return value, true
		}
	}
	return "", false
}

func (context *BaseContext) GetInterface(key string) (value interface{}, ok bool) {
	if value, ok := context.userData[key]; ok {
		return value, true
	}
	return nil, false
}

func (context *BaseContext) SetInt(key string, value int) {
	context.userData[key] = value
}

func (context *BaseContext) SetUInt(key string, value uint) {
	context.userData[key] = value
}

func (context *BaseContext) SetInt64(key string, value int64) {
	context.userData[key] = value
}

func (context *BaseContext) SetUInt64(key string, value uint64) {
	context.userData[key] = value
}

func (context *BaseContext) SetFloat(key string, value float64) {
	context.userData[key] = value
}

func (context *BaseContext) SetBool(key string, value bool) {
	context.userData[key] = value
}

func (context *BaseContext) SetString(key string, value string) {
	context.userData[key] = value
}

func (context *BaseContext) SetInterface(key string, value interface{}) {
	context.userData[key] = value
}
