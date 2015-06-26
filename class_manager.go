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
 * LastModified: Jun 26, 2015                             *
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
	mutex      sync.RWMutex
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
	cm.mutex.RLock()
	alias = cm.aliasCache[class]
	cm.mutex.RUnlock()
	return alias
}

// GetClass by alias.
func (cm *classManager) GetClass(alias string) (class reflect.Type) {
	cm.mutex.RLock()
	class = cm.classCache[alias]
	cm.mutex.RUnlock()
	return class
}

// GetTag by class.
func (cm *classManager) GetTag(class reflect.Type) (tag string) {
	cm.mutex.RLock()
	tag = cm.tagCache[class]
	cm.mutex.RUnlock()
	return tag
}

func initClassManager() *classManager {
	cm := new(classManager)
	cm.classCache = make(map[string]reflect.Type)
	cm.aliasCache = make(map[reflect.Type]string)
	cm.tagCache = make(map[reflect.Type]string)
	return cm
}

// ClassManager used to be register class with alias for hprose serialize/unserialize.
var ClassManager = initClassManager()
