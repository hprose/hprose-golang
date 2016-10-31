package main

import (
	"fmt"
	"math"
	"time"

	"github.com/hprose/hprose-golang/io"
)

func stringFloat64MapEncoder(w *io.Writer, v interface{}) {
	m := v.(map[string]float64)
	for key, val := range m {
		w.WriteString(key)
		w.WriteFloat(val, 64)
	}
}

func test(m map[string]float64) {
	start := time.Now()
	for i := 0; i < 500000; i++ {
		io.Marshal(m)
	}
	stop := time.Now()
	fmt.Println((stop.UnixNano() - start.UnixNano()) / 1000000)
}

func main() {
	m := make(map[string]float64)
	m["e"] = math.E
	m["pi"] = math.Pi
	m["ln2"] = math.Ln2
	test(m)
	io.RegisterMapEncoder((map[string]float64)(nil), stringFloat64MapEncoder)
	test(m)
}
