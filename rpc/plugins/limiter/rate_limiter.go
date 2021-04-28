/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/limiter/rate_limiter.go                      |
|                                                          |
| LastModified: Apr 29, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package limiter

import (
	"context"
	"math"
	"sync/atomic"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// RateLimiter plugin for hprose.
type RateLimiter struct {
	next             int64
	interval         float64
	permitsPerSecond int64
	maxPermits       float64
	timeout          time.Duration
}

// Option for RateLimiter.
type Option func(*RateLimiter)

// WithMaxPermits returns a maxPermits Option for RateLimiter.
func WithMaxPermits(maxPermits float64) Option {
	return func(l *RateLimiter) {
		l.maxPermits = maxPermits
	}
}

// WithTimeout returns a timeout Option for RateLimiter.
func WithTimeout(timeout time.Duration) Option {
	return func(l *RateLimiter) {
		l.timeout = timeout
	}
}

// NewRateLimiter returns a RateLimiter instance.
func NewRateLimiter(permitsPerSecond int64, options ...Option) *RateLimiter {
	l := &RateLimiter{
		next:             time.Now().UnixNano(),
		permitsPerSecond: permitsPerSecond,
		maxPermits:       math.Inf(0),
		timeout:          0,
		interval:         float64(time.Second) / float64(permitsPerSecond),
	}
	for _, option := range options {
		option(l)
	}
	return l
}

// Acquire is the core algorithm of RateLimiter.
func (l *RateLimiter) Acquire(ctx context.Context, tokens int) (err error) {
	now := time.Now().UnixNano()
	last := atomic.LoadInt64(&l.next)
	permits := float64(now-last)/l.interval - float64(tokens)
	if permits > l.maxPermits {
		permits = l.maxPermits
	}
	atomic.StoreInt64(&l.next, now-int64(permits*l.interval))
	if last <= now {
		return
	}
	delay := time.Duration(last - now)
	if l.timeout > 0 && delay > l.timeout {
		return core.ErrTimeout
	}
	ctx, cancel := context.WithTimeout(ctx, delay)
	<-ctx.Done()
	cancel()
	return
}

// IOHandler for RateLimiter.
func (l *RateLimiter) IOHandler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	if err = l.Acquire(ctx, len(request)); err != nil {
		return
	}
	return next(ctx, request)
}

// InvokeHandler for RateLimiter.
func (l *RateLimiter) InvokeHandler(ctx context.Context, name string, args []interface{}, next core.NextInvokeHandler) (result []interface{}, err error) {
	if err = l.Acquire(ctx, 1); err != nil {
		return
	}
	return next(ctx, name, args)
}

// PermitsPerSecond property of RateLimiter.
func (l *RateLimiter) PermitsPerSecond() int64 {
	return l.permitsPerSecond
}

// MaxPermits property of RateLimiter.
func (l *RateLimiter) MaxPermits() float64 {
	return l.maxPermits
}

// Timeout property of RateLimiter.
func (l *RateLimiter) Timeout() time.Duration {
	return l.timeout
}
