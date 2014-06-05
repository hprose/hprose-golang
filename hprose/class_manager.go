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
 * hprose/class_manager.go                                *
 *                                                        *
 * hprose ClassManager for Go.                            *
 *                                                        *
 * LastModified: Mar 27, 2014                             *
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
	tagCache   map[reflect.Type]string
	mutex      *sync.Mutex
}

// Register class with alias.
func (cm *classManager) Register(class reflect.Type, alias string, tag ...string) {
	cm.mutex.Lock()
	cm.classCache[alias] = class
	cm.aliasCache[class] = alias
	if len(tag) == 1 {
		cm.tagCache[class] = tag[0]
	}
	cm.mutex.Unlock()
}

// GetClassAlias by class.
func (cm *classManager) GetClassAlias(class reflect.Type) (alias string) {
	cm.mutex.Lock()
	alias = cm.aliasCache[class]
	cm.mutex.Unlock()
	return alias
}

// GetClass by alias.
func (cm *classManager) GetClass(alias string) (class reflect.Type) {
	cm.mutex.Lock()
	class = cm.classCache[alias]
	cm.mutex.Unlock()
	return class
}

// GetTag by class.
func (cm *classManager) GetTag(class reflect.Type) (tag string) {
	cm.mutex.Lock()
	tag = cm.tagCache[class]
	cm.mutex.Unlock()
	return tag
}

func initClassManager() *classManager {
	cm := classManager{make(map[string]reflect.Type), make(map[reflect.Type]string), make(map[reflect.Type]string), new(sync.Mutex)}
	return &cm
}

// ClassManager used to be register class with alias for hprose serialize/unserialize.
var ClassManager = initClassManager()
