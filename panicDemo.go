package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("Enter main func\n")
	defer func() {
		fmt.Printf("Enter defer func\n")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
	}()
	//recover的错误用法
	fmt.Printf("err use panic 1: %s\n", recover())
	panic(errors.New("something error"))

	p := recover()
	fmt.Printf("err use panic 2: %s\n", p)

	fmt.Printf("Exit main func")
}
