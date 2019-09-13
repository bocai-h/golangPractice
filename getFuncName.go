package main

import (
	"fmt"
	"runtime"
)

func getFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func addFunc(a int, b int) int {
	fmt.Printf("getFuncName()=%v\n", getFuncName())
	return a + b
}
func main() {
	addFunc(1, 2)
}
