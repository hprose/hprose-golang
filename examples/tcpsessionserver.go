package main

import "github.com/hprose/hprose-go"

var Session []map[string]interface{}

func GetSession(context hprose.Context) map[string]interface{} {
	sessionid, _ := context.GetInt("sessionid")
	return Session[sessionid]
}

type MyServerFilter struct{}

func (MyServerFilter) InputFilter(data []byte, context hprose.Context) []byte {
	if len(data) > 7 && data[0] == 's' && data[1] == 'i' && data[2] == 'd' {
		context.SetInt("sessionid", int(data[3])<<24|int(data[4])<<16|int(data[5])<<8|int(data[6]))
		data = data[7:]
	} else {
		context.SetInt("sessionid", len(Session))
		Session = append(Session, make(map[string]interface{}))
	}
	return data
}

func (MyServerFilter) OutputFilter(data []byte, context hprose.Context) []byte {
	sessionid, _ := context.GetInt("sessionid")
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

func inc(context hprose.Context) int {
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
	server.AddFilter(MyServerFilter{})
	server.AddFunction("inc", inc)
	server.Start()
}
