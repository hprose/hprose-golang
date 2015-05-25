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
 * hprose/1.1_compatible.go                               *
 *                                                        *
 * 1.1 compatible for Golang.                             *
 *                                                        *
 * LastModified: Feb 15, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"time"
)

// go 1.1 didn't have SetKeepAlivePeriod in TCPConn, so we should check it.
type iKeepAlivePeriod interface {
	SetKeepAlivePeriod(time.Duration) error
}
