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
 * rpc/basehttp_service.go                                *
 *                                                        *
 * hprose basehttp service for Go.                        *
 *                                                        *
 * LastModified: Nov 1, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

// BaseHTTPService is the hprose base http service
type BaseHTTPService struct {
	BaseService
	P3P                          bool
	GET                          bool
	CrossDomain                  bool
	AccessControlAllowOrigins    map[string]bool
	LastModified                 string
	Etag                         string
	crossDomainXMLFile           string
	crossDomainXMLContent        []byte
	clientAccessPolicyXMLFile    string
	clientAccessPolicyXMLContent []byte
}

// InitBaseHTTPService initializes BaseHTTPService
func (service *BaseHTTPService) InitBaseHTTPService() {
	t := time.Now().UTC()
	rand.Seed(t.UnixNano())
	service.InitBaseService()
	service.P3P = true
	service.GET = true
	service.CrossDomain = true
	service.AccessControlAllowOrigins = make(map[string]bool)
	service.LastModified = t.Format(time.RFC1123)
	service.Etag = `"` + strconv.FormatInt(rand.Int63(), 16) + `"`
}

// AddAccessControlAllowOrigin add access control allow origin
func (service *BaseHTTPService) AddAccessControlAllowOrigin(origins ...string) {
	for _, origin := range origins {
		service.AccessControlAllowOrigins[origin] = true
	}
}

// RemoveAccessControlAllowOrigin remove access control allow origin
func (service *BaseHTTPService) RemoveAccessControlAllowOrigin(origins ...string) {
	for _, origin := range origins {
		delete(service.AccessControlAllowOrigins, origin)
	}
}

// CrossDomainXMLFile return the cross domain xml file
func (service *BaseHTTPService) CrossDomainXMLFile() string {
	return service.crossDomainXMLFile
}

// CrossDomainXMLContent return the cross domain xml content
func (service *BaseHTTPService) CrossDomainXMLContent() []byte {
	return service.crossDomainXMLContent
}

// ClientAccessPolicyXMLFile return the client access policy xml file
func (service *BaseHTTPService) ClientAccessPolicyXMLFile() string {
	return service.clientAccessPolicyXMLFile
}

// ClientAccessPolicyXMLContent return the client access policy xml content
func (service *BaseHTTPService) ClientAccessPolicyXMLContent() []byte {
	return service.clientAccessPolicyXMLContent
}

// SetCrossDomainXMLFile set the cross domain xml file
func (service *BaseHTTPService) SetCrossDomainXMLFile(filename string) {
	service.crossDomainXMLFile = filename
	service.crossDomainXMLContent, _ = ioutil.ReadFile(filename)
}

// SetClientAccessPolicyXMLFile set the client access policy xml file
func (service *BaseHTTPService) SetClientAccessPolicyXMLFile(filename string) {
	service.clientAccessPolicyXMLFile = filename
	service.clientAccessPolicyXMLContent, _ = ioutil.ReadFile(filename)
}

// SetCrossDomainXMLContent set the cross domain xml content
func (service *BaseHTTPService) SetCrossDomainXMLContent(content []byte) {
	service.crossDomainXMLFile = ""
	service.crossDomainXMLContent = content
}

// SetClientAccessPolicyXMLContent set the client access policy xml content
func (service *BaseHTTPService) SetClientAccessPolicyXMLContent(content []byte) {
	service.clientAccessPolicyXMLFile = ""
	service.clientAccessPolicyXMLContent = content
}
