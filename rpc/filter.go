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
	filters  []Filter
	fmLocker sync.RWMutex
}

// Filter return the first filter
func (fm *filterManager) Filter() Filter {
	fm.fmLocker.RLock()
	defer fm.fmLocker.RUnlock()
	if len(fm.filters) == 0 {
		return nil
	}
	return fm.filters[0]
}

// NumFilter return the filter count
func (fm *filterManager) NumFilter() int {
	fm.fmLocker.RLock()
	defer fm.fmLocker.RUnlock()
	return len(fm.filters)
}

// FilterByIndex return the filter by index
func (fm *filterManager) FilterByIndex(index int) Filter {
	fm.fmLocker.RLock()
	defer fm.fmLocker.RUnlock()
	n := len(fm.filters)
	if index < 0 && index >= n {
		return nil
	}
	return fm.filters[index]
}

// SetFilter will replace the current filter settings
func (fm *filterManager) SetFilter(filter ...Filter) {
	fm.fmLocker.Lock()
	fm.filters = make([]Filter, len(filter))
	fm.AddFilter(filter...)
	fm.fmLocker.Unlock()
}

// AddFilter add the filter to this FilterManager
func (fm *filterManager) AddFilter(filter ...Filter) {
	fm.fmLocker.Lock()
	if len(filter) > 0 {
		fm.filters = append(fm.filters, filter...)
	}
	fm.fmLocker.Unlock()
}

// RemoveFilterByIndex remove the filter by the index
func (fm *filterManager) RemoveFilterByIndex(index int) {
	fm.fmLocker.Lock()
	n := len(fm.filters)
	if index < 0 && index >= n {
		fm.fmLocker.Unlock()
		return
	}
	if index == n-1 {
		fm.filters = fm.filters[:index]
	} else {
		fm.filters = append(fm.filters[:index], fm.filters[index+1:]...)
	}
	fm.fmLocker.Unlock()
}

func (fm *filterManager) removeFilter(filter Filter) {
	for i := range fm.filters {
		if fm.filters[i] == filter {
			fm.RemoveFilterByIndex(i)
			return
		}
	}
}

// RemoveFilter remove the filter from this FilterManager
func (fm *filterManager) RemoveFilter(filter ...Filter) {
	for i := range filter {
		fm.removeFilter(filter[i])
	}
}

func (fm *filterManager) inputFilter(data []byte, context Context) (out []byte, err error) {
	fm.fmLocker.RLock()
	defer func() {
		if e := recover(); e != nil {
			err = NewPanicError(e)
		}
		fm.fmLocker.RUnlock()
	}()
	for i := len(fm.filters) - 1; i >= 0; i-- {
		data = fm.filters[i].InputFilter(data, context)
	}
	out = data
	return
}

func (fm *filterManager) outputFilter(data []byte, context Context) (out []byte, err error) {
	fm.fmLocker.RLock()
	defer func() {
		if e := recover(); e != nil {
			err = NewPanicError(e)
		}
		defer fm.fmLocker.RUnlock()
	}()
	for i := range fm.filters {
		data = fm.filters[i].OutputFilter(data, context)
	}
	out = data
	return
}
