package main

import "fmt"

type printer func(output string)

func helloFunc(name string) {
	if name != "" {
		fmt.Printf("Hello, %s\n", name)
	} else {
		fmt.Printf("Hello, every one")
	}
}

func main() {
	var helloF printer
	helloF = helloFunc
	helloF("bocai")
}
