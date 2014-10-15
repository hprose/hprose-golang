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
	GetInt(key string) int
	GetUInt(key string) uint
	GetInt64(key string) int64
	GetUInt64(key string) uint64
	GetFloat(key string) float64
	GetBool(key string) bool
	GetString(key string) string
	GetInterface(key string) interface{}
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

func (context *BaseContext) GetInt(key string) int {
	return context.userData[key].(int)
}

func (context *BaseContext) GetUInt(key string) uint {
	return context.userData[key].(uint)
}

func (context *BaseContext) GetInt64(key string) int64 {
	return context.userData[key].(int64)
}

func (context *BaseContext) GetUInt64(key string) uint64 {
	return context.userData[key].(uint64)
}

func (context *BaseContext) GetFloat(key string) float64 {
	return context.userData[key].(float64)
}

func (context *BaseContext) GetBool(key string) bool {
	return context.userData[key].(bool)
}

func (context *BaseContext) GetString(key string) string {
	return context.userData[key].(string)
}

func (context *BaseContext) GetInterface(key string) interface{} {
	return context.userData[key]
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
