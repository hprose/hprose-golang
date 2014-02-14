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
 * LastModified: Feb 14, 2014                             *
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

// Register class with alias.
func (cm *classManager) Register(class reflect.Type, alias string) {
	cm.mutex.Lock()
	cm.classCache[alias] = class
	cm.aliasCache[class] = alias
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

func initClassManager() *classManager {
	cm := classManager{make(map[string]reflect.Type), make(map[reflect.Type]string), new(sync.Mutex)}
	return &cm
}

// ClassManager used to be register class with alias for hprose serialize/unserialize.
var ClassManager = initClassManager()
