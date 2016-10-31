package main

import (
	"fmt"
	"time"

	"github.com/hprose/hprose-golang/rpc"
)

type event struct{}

func (event) OnSubscribe(topic string, id string, service rpc.Service) {
	fmt.Println("client " + id + " subscribe topic: " + topic)
}

func (event) OnUnsubscribe(topic string, id string, service rpc.Service) {
	fmt.Println("client " + id + " unsubscribe topic: " + topic)
}

func main() {
	server := rpc.NewTCPServer("tcp4://0.0.0.0:2016/")
	server.Publish("time", 0, 0)
	server.Event = event{}
	var timer *time.Timer
	timer = time.AfterFunc(1*time.Second, func() {
		server.Broadcast("time", time.Now().String(), func(sended []string) {
			if len(sended) > 0 {
				fmt.Println(sended)
			}
		})
		timer.Reset(1 * time.Second)
	})
	server.Start()
}
