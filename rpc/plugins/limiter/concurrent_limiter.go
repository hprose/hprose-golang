/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/limiter/concurrent_limiter.go                |
|                                                          |
| LastModified: Mar 11, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package limiter

import (
	"context"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// ConcurrentLimiter plugin for hprose.
type ConcurrentLimiter struct {
	tasks                 chan struct{}
	maxConcurrentRequests int
	timeout               time.Duration
}

// NewConcurrentLimiter returns a ConcurrentLimiter instance.
func NewConcurrentLimiter(maxConcurrentRequests int, timeout ...time.Duration) *ConcurrentLimiter {
	t := time.Duration(0)
	if len(timeout) > 0 {
		t = timeout[0]
	}
	return &ConcurrentLimiter{
		tasks:                 make(chan struct{}, maxConcurrentRequests),
		maxConcurrentRequests: maxConcurrentRequests,
		timeout:               t,
	}
}

// Acquire returns immediately when the concurrentRequests is less than or equal to maxConcurrentRequests,
// otherwise it will block until timeout or any request are completed.
func (l *ConcurrentLimiter) Acquire() (err error) {
	if l.timeout > 0 {
		timer := time.NewTimer(l.timeout)
		select {
		case <-timer.C:
			return core.ErrTimeout
		case l.tasks <- struct{}{}:
			timer.Stop()
			return
		}
	}
	l.tasks <- struct{}{}
	return
}

// Release reqeust task queue.
func (l *ConcurrentLimiter) Release() {
	<-l.tasks
}

// Handler for ConcurrentLimiter.
func (l *ConcurrentLimiter) Handler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	if err = l.Acquire(); err != nil {
		return
	}
	defer l.Release()
	return next(ctx, request)
}

// ConcurrentRequests property of ConcurrentLimiter.
func (l *ConcurrentLimiter) ConcurrentRequests() int {
	return len(l.tasks)
}

// MaxConcurrentRequests property of ConcurrentLimiter.
func (l *ConcurrentLimiter) MaxConcurrentRequests() int {
	return l.maxConcurrentRequests
}

// Timeout property of ConcurrentLimiter.
func (l *ConcurrentLimiter) Timeout() time.Duration {
	return l.timeout
}
