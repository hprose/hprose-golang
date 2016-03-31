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

package hprose

import (
	"log"
	"time"

	"golang.org/x/net/context"
	etcd "github.com/coreos/etcd/client"
	"encoding/json"
	"math/rand"
	"strings"
)

type EtcdServersRepo struct {
	ServersRepo
	KeysAPI etcd.KeysAPI
}

func NewEtcdServersRepo(domain string, endpoints []string) *EtcdServersRepo {
	cfg := etcd.Config{
		Endpoints:               endpoints,
		Transport:               etcd.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := etcd.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}

	serversRepo := &EtcdServersRepo{
		KeysAPI: etcd.NewKeysAPI(etcdClient),
	}
	serversRepo.ServersMap = make(map[string]*Server)
	serversRepo.Domain = domain
	go serversRepo.WatchServers()
	return serversRepo
}

func (s *EtcdServersRepo) AddWorker(info *ServerInfo) {
	server := &Server{
		InGroup: true,
		ServerUrl:      info.ServerUrl,
		Domain:    info.Domain,
		UUID:        info.UUID,
		CPU:    info.CPU,
	}
	s.ServersMap[server.UUID] = server
}

func (s *EtcdServersRepo) GetPrimaryServer() (*Server) {
	return s.PrimaryServer
}

func (s *EtcdServersRepo) Update() {
	length := len(s.ServersMap)
	if length < 1 {
		s.retrieveAllServers()
	}
	length = len(s.ServersMap)
	if length < 1 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(length)
	count := 0
	for _, server := range s.ServersMap {
		if count == index {
			s.PrimaryServer = server
			//println("EtcdServersRepo.Update", server.ServerUrl)
			return
		}
		count++
	}
}

func (s *EtcdServersRepo) retrieveAllServers() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()

	rsp, err := s.KeysAPI.Get(ctx, "hprose-service/" + s.Domain, &etcd.GetOptions{Recursive: true, Sort: true})
	if err != nil && !strings.HasPrefix(err.Error(), "100: Key not found") {
		log.Println("Error retrieve hprose servers:", err)
		return
	}
	if rsp == nil {
		log.Println("Retrieve none hprose servers:")
		return
	}

	for _, node := range rsp.Node.Nodes {
		info := &ServerInfo{}
		err := json.Unmarshal([]byte(node.Value), info)

		if err != nil {
			log.Println("Error retrieve hprose servers during unmarshall the response:", err, node.Value)
		}
		if _, ok := s.ServersMap[info.UUID]; ok {
			s.UpdateServer(info)
		} else {
			s.AddWorker(info)
		}
	}

/*	if rsp.Node.Nodes.Len() >= 1 {
		println("EtcdServersRepo", rsp.Node.Nodes[0].Value)
		for _, node := range rsp.Node.Nodes[0].Nodes {
			info := &ServerInfo{}
			err := json.Unmarshal([]byte(node.Value), info)

			if err != nil {
				log.Println("Error retrieve hprose servers during unmarshall the response:", err, node.Value)
			}
			if _, ok := s.ServersMap[info.UUID]; ok {
				s.UpdateServer(info)
			} else {
				s.AddWorker(info)
			}
		}
	}*/

}

func (s *EtcdServersRepo) UpdateServer(info *ServerInfo) {
	server := s.ServersMap[info.UUID]
	server.InGroup = true
}

func (s *EtcdServersRepo) WatchServers() {
	api := s.KeysAPI
	watcher := api.Watcher("hprose-service/" + s.Domain, &etcd.WatcherOptions{
		Recursive: true,
	})
	for {
		res, err := watcher.Next(context.Background())
		if err != nil {
			log.Println("Error watch hprose servers:", err)
			break
		}
		if res.Action == "expire" {
			if res.Node.Key == "" && strings.Index(res.Node.Key, "/") <0{
				return
			}
			spliteKeys := strings.Split(res.Node.Key, "/")
			counts := len(spliteKeys)
			key := spliteKeys[counts-1]

			server, ok := s.ServersMap[key]
			if ok {
				server.InGroup = false
				delete(s.ServersMap, server.UUID)
				s.Update()
			}
		} else if res.Action == "set" || res.Action == "update" {
			info := &ServerInfo{}
			err := json.Unmarshal([]byte(res.Node.Value), info)
			if err != nil {
				log.Print(err)
			}
			if _, ok := s.ServersMap[info.UUID]; ok {
				s.UpdateServer(info)
			} else {
				s.AddWorker(info)
				s.Update()

			}
		} else if res.Action == "delete" {
			delete(s.ServersMap, res.Node.Key)
			s.Update()
		}
	}

}

