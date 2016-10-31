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
 * LastModified: Oct 2, 2016                              *
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

type baseHTTPService struct {
	baseService
	P3P                          bool
	GET                          bool
	CrossDomain                  bool
	accessControlAllowOrigins    map[string]bool
	lastModified                 string
	etag                         string
	crossDomainXMLFile           string
	crossDomainXMLContent        []byte
	clientAccessPolicyXMLFile    string
	clientAccessPolicyXMLContent []byte
}

func (service *baseHTTPService) initBaseHTTPService() {
	t := time.Now().UTC()
	rand.Seed(t.UnixNano())
	service.initBaseService()
	service.P3P = true
	service.GET = true
	service.CrossDomain = true
	service.accessControlAllowOrigins = make(map[string]bool)
	service.lastModified = t.Format(time.RFC1123)
	service.etag = `"` + strconv.FormatInt(rand.Int63(), 16) + `"`
}

// AddAccessControlAllowOrigin add access control allow origin
func (service *baseHTTPService) AddAccessControlAllowOrigin(origins ...string) {
	for _, origin := range origins {
		service.accessControlAllowOrigins[origin] = true
	}
}

// RemoveAccessControlAllowOrigin remove access control allow origin
func (service *baseHTTPService) RemoveAccessControlAllowOrigin(origins ...string) {
	for _, origin := range origins {
		delete(service.accessControlAllowOrigins, origin)
	}
}

// CrossDomainXMLFile return the cross domain xml file
func (service *baseHTTPService) CrossDomainXMLFile() string {
	return service.crossDomainXMLFile
}

// CrossDomainXMLContent return the cross domain xml content
func (service *baseHTTPService) CrossDomainXMLContent() []byte {
	return service.crossDomainXMLContent
}

// ClientAccessPolicyXMLFile return the client access policy xml file
func (service *baseHTTPService) ClientAccessPolicyXMLFile() string {
	return service.clientAccessPolicyXMLFile
}

// ClientAccessPolicyXMLContent return the client access policy xml content
func (service *baseHTTPService) ClientAccessPolicyXMLContent() []byte {
	return service.clientAccessPolicyXMLContent
}

// SetCrossDomainXMLFile set the cross domain xml file
func (service *baseHTTPService) SetCrossDomainXMLFile(filename string) {
	service.crossDomainXMLFile = filename
	service.crossDomainXMLContent, _ = ioutil.ReadFile(filename)
}

// SetClientAccessPolicyXMLFile set the client access policy xml file
func (service *baseHTTPService) SetClientAccessPolicyXMLFile(filename string) {
	service.clientAccessPolicyXMLFile = filename
	service.clientAccessPolicyXMLContent, _ = ioutil.ReadFile(filename)
}

// SetCrossDomainXMLContent set the cross domain xml content
func (service *baseHTTPService) SetCrossDomainXMLContent(content []byte) {
	service.crossDomainXMLFile = ""
	service.crossDomainXMLContent = content
}

// SetClientAccessPolicyXMLContent set the client access policy xml content
func (service *baseHTTPService) SetClientAccessPolicyXMLContent(content []byte) {
	service.clientAccessPolicyXMLFile = ""
	service.clientAccessPolicyXMLContent = content
}
