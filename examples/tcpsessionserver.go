package main

import (
	"github.com/hprose/hprose-go/hprose"
	"os"
)

var SessionIdMap = make(map[interface{}]int)
var Session []map[string]interface{}

func GetSession(context interface{}) map[string]interface{} {
	return Session[SessionIdMap[context]]
}

type MyServerFilter struct{}

func (MyServerFilter) InputFilter(data []byte, context interface{}) []byte {
	if len(data) > 7 && data[0] == 's' && data[1] == 'i' && data[2] == 'd' {
		sessionid := int(data[3])<<24 | int(data[4])<<16 | int(data[5])<<8 | int(data[6])
		SessionIdMap[context] = sessionid
		data = data[7:]
	} else {
		sessionid := len(Session)
		SessionIdMap[context] = sessionid
		Session = append(Session, make(map[string]interface{}))
	}
	return data
}

func (MyServerFilter) OutputFilter(data []byte, context interface{}) []byte {
	sessionid := SessionIdMap[context]
	buf := make([]byte, 7+len(data))
	buf[0] = 's'
	buf[1] = 'i'
	buf[2] = 'd'
	buf[3] = byte(sessionid >> 24 & 0xff)
	buf[4] = byte(sessionid >> 16 & 0xff)
	buf[5] = byte(sessionid >> 8 & 0xff)
	buf[6] = byte(sessionid & 0xff)
	copy(buf[7:], data)
	return buf
}

func inc(context interface{}) int {
	session := GetSession(context)
	n, ok := session["n"]
	if !ok {
		session["n"] = 0
		return 0
	}
	i := n.(int) + 1
	session["n"] = i
	return i
}

func main() {
	server := hprose.NewTcpServer("tcp4://:4321/")
	server.Filters = append(server.Filters, MyServerFilter{})
	server.AddFunction("inc", inc)
	server.ThreadCount = 16
	server.Start()
	b := make([]byte, 1)
	os.Stdin.Read(b)
	server.Stop()
}
