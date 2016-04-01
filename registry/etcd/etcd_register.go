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
 * Author: Henry Hu <henry.pf.hu@gmail.com                *
 *                                                        *
\**********************************************************/

package etcd

import (
	"runtime"
	"encoding/json"
	"time"
	"log"
	"net"
	"os"
	"os/signal"
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/net/context"
	etcd "github.com/coreos/etcd/client"

//"crypto/rand"*
)

type EtcdRegister struct {
	UUID      string
	Domain    string
	ServerUrl string
	KeysAPI   etcd.KeysAPI
	signal    chan os.Signal
}

// ServerInfo is the service register information to etcd
type ServerInfo struct {
	UUID      string
	Domain    string
	ServerUrl string
	CPU       int
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func RegisterServer(domain, serverUrl string, etcEndpoints []string) {
	etcdCfg := etcd.Config{
		Endpoints:               etcEndpoints,
		Transport:               etcd.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := etcd.New(etcdCfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}

	e := &EtcdRegister{
		UUID: uuid(serverUrl),
		Domain: domain,
		ServerUrl: serverUrl,
		KeysAPI: etcd.NewKeysAPI(etcdClient),
	}

	go e.HeartBeat()

	e.signal = make(chan os.Signal, 1)
	signal.Notify(e.signal, os.Interrupt, os.Kill)
}

func (e *EtcdRegister) HeartBeat() {
	for {
		select {
		case <-e.signal:
			return
		case <-time.After(time.Second * 3):
			e.updateHB()
		}
	}
}

// Stop the hprose tcp server
func (e *EtcdRegister) updateHB() {
	api := e.KeysAPI

	info := &ServerInfo{
		UUID: e.UUID,
		Domain: e.Domain,
		ServerUrl:   e.ServerUrl,
		CPU:  runtime.NumCPU(),
	}

	key := "hprose-service/" + info.Domain + "/" + info.UUID
	value, _ := json.Marshal(info)

	_, err := api.Set(context.Background(), key, string(value), &etcd.SetOptions{
		TTL: time.Second * 10,
	})
	if err != nil {
		log.Println("Error update ServerInfo:", err)
	}
}

func uuid(serverUrl string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(serverUrl))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
	/*	b := make([]byte, 16)
		rand.Read(b)
		b[6] = (b[6] & 0x0f) | 0x40
		b[8] = (b[8] & 0x3f) | 0x80
		return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])*/
}