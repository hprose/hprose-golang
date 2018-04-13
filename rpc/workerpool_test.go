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
 * rpc/workerpool_test.go                                 *
 *                                                        *
 * hprose gorouine pool test for Go.                      *
 *                                                        *
 * LastModified: Apr 13, 2018                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	wp := WorkerPool{}
	wp.Start()
	defer wp.Stop()

	for i := 0; i < 10; i++ {
		wp.Go(func() {
			fmt.Println("hello")
		})
	}
}

func TestStop(t *testing.T) {
	wp := WorkerPool{}
	wp.Start()
	defer wp.Stop()

	for i := 0; i < 10; i++ {
		wp.Go(func() {
			fmt.Println("hello")
		})
	}
}

func BenchmarkPool(b *testing.B) {
	wp := WorkerPool{}
	wp.Start()
	defer wp.Stop()
	b.ResetTimer()

	for k := 0; k < b.N; k++ {
		wp.Go(func() {
			rand.Int()
			time.Sleep(1 * time.Microsecond)
		})
	}

}

func BenchmarkGoroutine(b *testing.B) {
	b.ResetTimer()

	for k := 0; k < b.N; k++ {
		go func() {
			rand.Int()
			time.Sleep(1 * time.Microsecond)
		}()
	}
}