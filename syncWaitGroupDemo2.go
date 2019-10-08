package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	coordinateWithWaitGroup()
}
func coordinateWithWaitGroup() {
	total := 12
	stride := 3
	var num int32
	fmt.Printf("The num: %d [with sync.WaitGroup]\n", num)
	var wg sync.WaitGroup
	for i := 1; i < total; i = i + stride {
		wg.Add(stride)
		for j := 0; j < stride; j++ {
			go addNum(&num, i+j, wg.Done)
		}
		wg.Wait()
		fmt.Printf("#######next wait [i=%d]\n", i)
	}
	fmt.Println("End.")
}
func addNum(numP *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		newNum := currNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			break
		} else {
			fmt.Printf("The CAS operation failed.[%d-%d]\n", id, i)
		}
	}
}
