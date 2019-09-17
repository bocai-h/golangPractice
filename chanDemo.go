package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := make([]int, 0, 10)
	fmt.Printf("unsafe.Sizeof(s1)=%v\n", unsafe.Sizeof(s1))
}
