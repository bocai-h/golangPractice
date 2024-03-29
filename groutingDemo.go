package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var count uint32

var trigger = func(i uint32, fn func()) {
	for {
		if n := atomic.LoadUint32(&count); n == i {
			fn()
			atomic.AddUint32(&count, 1)
			break
		}
		time.Sleep(time.Nanosecond)
	}
}

func main() {
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {})
}
