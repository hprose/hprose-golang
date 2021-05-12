/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/circuitbreaker/circuitbreaker.go             |
|                                                          |
| LastModified: May 12, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package circuitbreaker

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// ErrBreaker for circuit breaker.
var ErrBreaker = errors.New("service breaked")

// MockService for circuit breaker.
type MockService = func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error)

// CircuitBreaker plugin for hprose.
type CircuitBreaker struct {
	lastFailTime int64
	failCount    uint64
	threshold    uint64
	recoverTime  time.Duration
	mockService  MockService
}

// Option for CircuitBreaker.
type Option func(*CircuitBreaker)

// WithThreshold returns a threshold Option for CircuitBreaker.
func WithThreshold(threshold uint64) Option {
	return func(cb *CircuitBreaker) {
		cb.threshold = threshold
	}
}

// WithRecoverTime returns a recoverTime Option for CircuitBreaker.
func WithRecoverTime(recoverTime time.Duration) Option {
	return func(cb *CircuitBreaker) {
		cb.recoverTime = recoverTime
	}
}

// WithMockService returns a mockService Option for CircuitBreaker.
func WithMockService(mockService MockService) Option {
	return func(cb *CircuitBreaker) {
		cb.mockService = mockService
	}
}

// New returns a CircuitBreaker instance.
func New(options ...Option) *CircuitBreaker {
	cb := &CircuitBreaker{
		threshold:   5,
		recoverTime: time.Second * 30,
	}
	for _, option := range options {
		option(cb)
	}
	return cb
}

// Threshold property of CircuitBreaker.
func (cb *CircuitBreaker) Threshold() uint64 {
	return cb.threshold
}

// RecoverTime property of CircuitBreaker.
func (cb *CircuitBreaker) RecoverTime() time.Duration {
	return cb.recoverTime
}

// MockService property of CircuitBreaker.
func (cb *CircuitBreaker) MockService() MockService {
	return cb.mockService
}

// IOHandler for CircuitBreaker.
func (cb *CircuitBreaker) IOHandler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	if atomic.LoadUint64(&cb.failCount) > cb.threshold {
		interval := time.Duration(time.Now().UnixNano() - atomic.LoadInt64(&cb.lastFailTime))
		if interval < cb.recoverTime {
			return nil, ErrBreaker
		}
		atomic.StoreUint64(&cb.failCount, cb.threshold>>1)
	}
	defer func() {
		if e := recover(); e != nil {
			err = core.NewPanicError(e)
		}
		if err != nil {
			atomic.AddUint64(&cb.failCount, 1)
			atomic.StoreInt64(&cb.lastFailTime, time.Now().UnixNano())
		}
	}()
	response, err = next(ctx, request)
	if err == nil {
		atomic.StoreUint64(&cb.failCount, 0)
	}
	return
}

// InvokeHandler for CircuitBreaker.
func (cb *CircuitBreaker) InvokeHandler(ctx context.Context, name string, args []interface{}, next core.NextInvokeHandler) (result []interface{}, err error) {
	if cb.mockService == nil {
		return next(ctx, name, args)
	}
	if result, err = next(ctx, name, args); err == ErrBreaker {
		return cb.mockService(ctx, name, args)
	}
	return
}
