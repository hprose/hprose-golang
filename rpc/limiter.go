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
 * LastModified: Oct 2, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import "sync"

type limiter struct {
	cond                  sync.Cond
	requestCount          int
	MaxConcurrentRequests int
}

func (limiter *limiter) initLimiter() {
	limiter.MaxConcurrentRequests = 10
	limiter.cond.L = &sync.Mutex{}
}

func (limiter *limiter) limit() {
	for {
		if limiter.requestCount < limiter.MaxConcurrentRequests {
			break
		}
		limiter.cond.Wait()
	}
	limiter.requestCount++
}

func (limiter *limiter) unlimit() {
	limiter.requestCount--
	limiter.cond.Signal()
}

func (limiter *limiter) reset() {
	limiter.requestCount = 0
	for i := limiter.MaxConcurrentRequests; i > 0; i-- {
		limiter.cond.Signal()
	}
}
