/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/http/cookie.go                                       |
|                                                          |
| LastModified: Mar 6, 2022                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package cookie

type CookieManagerOption int

const (
	NoCookieManager CookieManagerOption = iota
	GlobalCookieManager
	ClientCookieManager
)
