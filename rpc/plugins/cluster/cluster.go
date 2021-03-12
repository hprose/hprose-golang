/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/cluster/cluster.go                           |
|                                                          |
| LastModified: Mar 6, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package cluster

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// Config for cluster.
type Config struct {
	Retry       int
	Idempotent  bool
	OnSuccess   func(context.Context)
	OnFailure   func(context.Context)
	OnRetry     func(context.Context) time.Duration
	minInterval time.Duration
	maxInterval time.Duration
}

// Option for cluster config.
type Option func(*Config)

// WithRetry returns a retry Option for cluster config.
func WithRetry(retry int) Option {
	return func(c *Config) {
		c.Retry = retry
	}
}

// WithIdempotent returns an idempotent Option for cluster config.
func WithIdempotent(idempotent bool) Option {
	return func(c *Config) {
		c.Idempotent = idempotent
	}
}

// WithMinInterval returns a retry minInterval Option for cluster config.
func WithMinInterval(minInterval time.Duration) Option {
	return func(c *Config) {
		c.minInterval = minInterval
	}
}

// WithMaxInterval returns a retry maxInterval Option for cluster config.
func WithMaxInterval(maxInterval time.Duration) Option {
	return func(c *Config) {
		c.maxInterval = maxInterval
	}
}

func getIndex(index *int64, n int64) int64 {
	if n > 1 {
		if i := atomic.AddInt64(index, 1); i < n {
			return i
		}
		atomic.StoreInt64(index, 0)
	}
	return 0
}

// FailoverConfig for cluster.
func FailoverConfig(options ...Option) (config Config) {
	config.Retry = 10
	config.minInterval = time.Millisecond * 500
	config.maxInterval = time.Second * 5
	for _, option := range options {
		option(&config)
	}
	var index int64
	config.OnFailure = func(ctx context.Context) {
		clientContext := core.GetClientContext(ctx)
		urls := clientContext.Client().URLs
		clientContext.URL = urls[getIndex(&index, int64(len(urls)))]
	}
	config.OnRetry = func(ctx context.Context) time.Duration {
		clientContext := core.GetClientContext(ctx)
		retried := clientContext.Items().GetInt("retried") + 1
		clientContext.Items().Set("retried", retried)
		interval := config.minInterval * time.Duration(retried-len(clientContext.Client().URLs))
		if interval > config.maxInterval {
			interval = config.maxInterval
		}
		return interval
	}
	return
}

// FailtryConfig for cluster.
func FailtryConfig(options ...Option) (config Config) {
	config.Retry = 10
	config.minInterval = time.Millisecond * 500
	config.maxInterval = time.Second * 5
	for _, option := range options {
		option(&config)
	}
	config.OnRetry = func(ctx context.Context) time.Duration {
		clientContext := core.GetClientContext(ctx)
		retried := clientContext.Items().GetInt("retried") + 1
		clientContext.Items().Set("retried", retried)
		interval := config.minInterval * time.Duration(retried)
		if interval > config.maxInterval {
			interval = config.maxInterval
		}
		return interval
	}
	return
}

// FailfastConfig for cluster.
func FailfastConfig(onFailure func(context.Context)) (config Config) {
	config.Retry = 0
	config.OnFailure = onFailure
	return
}

// Cluster plugin for hprose.
type Cluster struct {
	Config
}

// New returns a Cluster instance.
func New(config ...Config) (cluster *Cluster) {
	cluster = &Cluster{}
	if len(config) > 0 {
		cluster.Config = config[0]
		if cluster.Retry < 0 {
			cluster.Retry = 10
		}
	} else {
		cluster.Config = FailoverConfig()
	}
	return
}

// Handler for Cluster.
func (c *Cluster) Handler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = core.NewPanicError(e)
		}
		if err == nil {
			return
		}
		if c.OnFailure != nil {
			c.OnFailure(ctx)
		}
		if c.OnRetry == nil {
			return
		}
		clientContext := core.GetClientContext(ctx)
		idempotent := clientContext.Items().GetBool("idempotent", c.Idempotent)
		retry := clientContext.Items().GetInt("retry", c.Retry)
		retried := clientContext.Items().GetInt("retried")
		if idempotent && (retried < retry) {
			interval := c.OnRetry(ctx)
			if interval > 0 {
				time.Sleep(interval)
			}
			response, err = c.Handler(ctx, request, next)
		}
	}()
	if response, err = next(ctx, request); err == nil && c.OnSuccess != nil {
		c.OnSuccess(ctx)
	}
	return
}

// Forking on Cluster.
func Forking(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	clientContext := core.GetClientContext(ctx)
	urls := clientContext.Client().URLs
	n := len(urls)
	if n == 0 {
		return next(ctx, request)
	}
	count := int64(n)
	var once sync.Once
	done := make(chan struct{})
	for i := 0; i < n; i++ {
		forkingContext := clientContext.Clone().(*core.ClientContext)
		forkingContext.URL = urls[i]
		go func(ctx context.Context) {
			defer func() {
				if e := recover(); e != nil && atomic.AddInt64(&count, -1) <= 0 {
					once.Do(func() {
						err = core.NewPanicError(e)
						close(done)
					})
				}
			}()
			resp, e := next(ctx, request)
			if e == nil {
				once.Do(func() {
					response = resp
					close(done)
				})
				return
			}
			if atomic.AddInt64(&count, -1) <= 0 {
				once.Do(func() {
					err = e
					close(done)
				})
			}
		}(core.WithContext(ctx, forkingContext))
	}
	<-done
	return
}

// Broadcast on Cluster.
func Broadcast(ctx context.Context, name string, args []interface{}, next core.NextInvokeHandler) (result []interface{}, err error) {
	clientContext := core.GetClientContext(ctx)
	urls := clientContext.Client().URLs
	n := len(urls)
	if n == 0 {
		return next(ctx, name, args)
	}
	var wg sync.WaitGroup
	var once sync.Once
	result = make([]interface{}, n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		forkingContext := clientContext.Clone().(*core.ClientContext)
		forkingContext.URL = urls[i]
		go func(i int, ctx context.Context) {
			defer func() {
				if e := recover(); e != nil {
					once.Do(func() { err = core.NewPanicError(e) })
				}
				wg.Done()
			}()
			var e error
			if result[i], e = next(ctx, name, args); e != nil {
				once.Do(func() { err = e })
			}
		}(i, core.WithContext(ctx, forkingContext))
	}
	wg.Wait()
	return
}
