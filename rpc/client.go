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
 * rpc/client.go                                          *
 *                                                        *
 * hprose rpc client for Go.                              *
 *                                                        *
 * LastModified: Oct 27, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"crypto/tls"
	"errors"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"
)

// InvokeSettings is the invoke settings of hprose client
type InvokeSettings struct {
	ByRef          bool
	Simple         bool
	Idempotent     bool
	Failswitch     bool
	Oneway         bool
	JSONCompatible bool
	Retry          int
	Mode           ResultMode
	Timeout        time.Duration
	ResultTypes    []reflect.Type
	userData       map[string]interface{}
}

// SetUserData on InvokeSettings
func (settings *InvokeSettings) SetUserData(data map[string]interface{}) {
	settings.userData = data
}

// Callback is the callback function type of Client.Go
type Callback func([]reflect.Value, error)

// Client is hprose client
type Client interface {
	URL() *url.URL
	URI() string
	SetURI(uri string)
	URIList() []string
	SetURIList(uriList []string)
	TLSClientConfig() *tls.Config
	SetTLSClientConfig(config *tls.Config)
	Retry() int
	SetRetry(value int)
	Timeout() time.Duration
	SetTimeout(value time.Duration)
	Failround() int
	SetEvent(ClientEvent)
	Filter() Filter
	FilterByIndex(index int) Filter
	SetFilter(filter ...Filter) Client
	AddFilter(filter ...Filter) Client
	RemoveFilterByIndex(index int) Client
	RemoveFilter(filter ...Filter) Client
	AddInvokeHandler(handler ...InvokeHandler) Client
	AddBeforeFilterHandler(handler ...FilterHandler) Client
	AddAfterFilterHandler(handler ...FilterHandler) Client
	SetUserData(userdata map[string]interface{}) Client
	UseService(remoteService interface{}, namespace ...string)
	Invoke(string, []reflect.Value, *InvokeSettings) ([]reflect.Value, error)
	Go(string, []reflect.Value, *InvokeSettings, Callback)
	Close()
	AutoID() (string, error)
	ID() string
	Subscribe(name string, id string, settings *InvokeSettings, callback interface{}) (err error)
	Unsubscribe(name string, id ...string)
	IsSubscribed(name string) bool
	SubscribedList() []string
}

// ClientContext is the hprose client context
type ClientContext struct {
	baseContext
	InvokeSettings
	Retried int
	Client  Client
}

// NewClient is the constructor of Client
func NewClient(uri ...string) Client {
	return clientFactories[checkAddresses(uri, allSchemes)](uri...)
}

// UseFastHTTPClient as the default http client
func UseFastHTTPClient() {
	registerClientFactory("http", newFastHTTPClient)
	registerClientFactory("https", newFastHTTPClient)
}

var httpSchemes = []string{"http", "https"}
var tcpSchemes = []string{"tcp", "tcp4", "tcp6"}
var unixSchemes = []string{"unix"}
var websocketSchemes = []string{"ws", "wss"}
var allSchemes = []string{"http", "https", "tcp", "tcp4", "tcp6", "unix", "ws", "wss"}

func checkAddresses(uriList []string, schemes []string) (scheme string) {
	count := len(uriList)
	if count < 1 {
		panic(errURIListEmpty)
	}
	u, err := url.Parse(uriList[0])
	if err != nil {
		panic(err)
	}
	scheme = u.Scheme
	if sort.SearchStrings(schemes, scheme) == len(schemes) {
		panic(errors.New("This client desn't support " + scheme + " scheme."))
	}
	for i := 1; i < count; i++ {
		u, err := url.Parse(uriList[i])
		if err != nil {
			panic(err)
		}
		if scheme != u.Scheme {
			panic(errNotSupportMultpleProtocol)
		}
	}
	return
}

var clientFactories = make(map[string]func(...string) Client)

func registerClientFactory(scheme string, newClient func(...string) Client) {
	clientFactories[strings.ToLower(scheme)] = newClient
}

func init() {
	registerClientFactory("http", newHTTPClient)
	registerClientFactory("https", newHTTPClient)
	registerClientFactory("tcp", newTCPClient)
	registerClientFactory("tcp4", newTCPClient)
	registerClientFactory("tcp6", newTCPClient)
	registerClientFactory("unix", newUnixClient)
	registerClientFactory("ws", newWebSocketClient)
	registerClientFactory("wss", newWebSocketClient)
}
