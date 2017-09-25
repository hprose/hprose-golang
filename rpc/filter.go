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
 * rpc/filter.go                                          *
 *                                                        *
 * hprose filter interface for Go.                        *
 *                                                        *
 * LastModified: May 22, 2017                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import "sync"

// Filter is hprose filter
type Filter interface {
	InputFilter(data []byte, context Context) []byte
	OutputFilter(data []byte, context Context) []byte
}

// filterManager is the filter manager
type filterManager struct {
	filters []Filter
	sync.RWMutex
}

// Filter return the first filter
func (fm *filterManager) FirstFilter() (f Filter) {
	fm.Lock()
	if len(fm.filters) > 0 {
		f = fm.filters[0]
	}
	fm.Unlock()
	return
}

// NumFilter return the filter count
func (fm *filterManager) NumFilter() int {
	return len(fm.filters)
}

// FilterByIndex return the filter by index
func (fm *filterManager) FilterByIndex(index int) (f Filter) {
	fm.Lock()
	if index >= 0 && index < len(fm.filters) {
		f = fm.filters[index]
	}
	fm.Unlock()
	return
}

// SetFilter will replace the current filter settings
func (fm *filterManager) SetFilters(filters ...Filter) {
	fm.Lock()
	fm.filters = make([]Filter, 0, len(filters))
	fm.filters = append(fm.filters, filters...)
	// it will be deadlock!!!
	//fm.AddFilter(filter...)
	fm.Unlock()
}

// AddFilters add the filter to this FilterManager
func (fm *filterManager) AddFilters(filters ...Filter) {
	fm.Lock()
	fm.filters = append(fm.filters, filters...)
	fm.Unlock()
}

// RemoveFilterByIndex remove the filter by the index
func (fm *filterManager) RemoveFilterByIndex(index int) {
	fm.Lock()
	if 0 > index || index >= len(fm.filters) {
		return
	}
	fm.filters = append(fm.filters[:index], fm.filters[index+1:]...)
	fm.Unlock()
	return
}

func (fm *filterManager) removeFilter(filter Filter) {
	// if goroutines come concurrent to modify filters.
	// it may be delete index 0 filter , it actually deletes index 1 filter.
	fm.Lock()
	for i := range fm.filters {
		if fm.filters[i] == filter {
			fm.filters = append(fm.filters[:i], fm.filters[i+1:]...)
			return
		}
	}
	fm.Unlock()
}

// RemoveFilter remove the filter from this FilterManager
func (fm *filterManager) RemoveFilters(filters ...Filter) {
	fm.Lock()
	for _, filter := range filters {
		for i := range fm.filters {
			if fm.filters[i] == filter {
				fm.filters = append(fm.filters[:i], fm.filters[i+1:]...)
				break
			}
		}
	}
	fm.Unlock()
}

func (fm *filterManager) inputFilter(data []byte, context Context) (out []byte, err error) {
	fm.RLock()
	defer func() {
		if e := recover(); e != nil {
			err = NewPanicError(e)
		}
		fm.RUnlock()
	}()
	// don't change original data
	out = make([]byte, len(data))
	copy(out, data)
	for i := len(fm.filters) - 1; i >= 0; i-- {
		out = fm.filters[i].InputFilter(out, context)
	}
	return
}

func (fm *filterManager) outputFilter(data []byte, context Context) (out []byte, err error) {
	fm.RLock()
	defer func() {
		if e := recover(); e != nil {
			err = NewPanicError(e)
		}
		fm.RUnlock()
	}()
	// don't change original data
	out = make([]byte, len(data))
	copy(out, data)
	for i := range fm.filters {
		out = fm.filters[i].OutputFilter(out, context)
	}
	return
}
