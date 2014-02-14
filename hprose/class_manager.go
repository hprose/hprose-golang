/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.net/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * hprose/class_manager.go                                *
 *                                                        *
 * hprose ClassManager for Go.                            *
 *                                                        *
 * LastModified: Jan 26, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"reflect"
	"sync"
)

type classManager struct {
	classCache map[string]reflect.Type
	aliasCache map[reflect.Type]string
	mutex      *sync.Mutex
}

type IClassManager interface {
	Register(reflect.Type, string)
	GetClassAlias(reflect.Type) string
	GetClass(string) reflect.Type
}

func (self *classManager) Register(class reflect.Type, alias string) {
	self.mutex.Lock()
	self.classCache[alias] = class
	self.aliasCache[class] = alias
	self.mutex.Unlock()
}

func (self *classManager) GetClassAlias(class reflect.Type) (alias string) {
	self.mutex.Lock()
	alias = self.aliasCache[class]
	self.mutex.Unlock()
	return alias
}

func (self *classManager) GetClass(alias string) (class reflect.Type) {
	self.mutex.Lock()
	class = self.classCache[alias]
	self.mutex.Unlock()
	return class
}

func initClassManager() *classManager {
	cm := classManager{make(map[string]reflect.Type), make(map[reflect.Type]string), new(sync.Mutex)}
	return &cm
}

var ClassManager IClassManager = initClassManager()
