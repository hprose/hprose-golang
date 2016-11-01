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
 * rpc/limiter.go                                         *
 *                                                        *
 * hprose client requests limiter for Go.                 *
 *                                                        *
 * LastModified: Nov 1, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import "sync"

// Limiter is a request limiter
type Limiter struct {
	MaxConcurrentRequests int
	requestCount          int
	sync.Cond
}

// InitLimiter initializes Limiter
func (limiter *Limiter) InitLimiter() {
	limiter.MaxConcurrentRequests = 10
	limiter.L = &sync.Mutex{}
}

// Limit the request
func (limiter *Limiter) Limit() {
	for {
		if limiter.requestCount < limiter.MaxConcurrentRequests {
			break
		}
		limiter.Wait()
	}
	limiter.requestCount++
}

// Unlimit the request
func (limiter *Limiter) Unlimit() {
	limiter.requestCount--
	limiter.Signal()
}

// Reset the Limiter
func (limiter *Limiter) Reset() {
	limiter.requestCount = 0
	for i := limiter.MaxConcurrentRequests; i > 0; i-- {
		limiter.Signal()
	}
}
