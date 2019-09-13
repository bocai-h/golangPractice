package main

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

type opFunc func(a int, b int) int

// gen a high order function
func genCalculate(a int, b int, opfunc opFunc) (rFunc opFunc, err error) {
	if opfunc == nil {
		err = errors.New("opfunc is nil")
		return
	}
	rFunc = func(i int, j int) int {
		return opfunc(i, j)
	}
	return
}

func main() {
	addFunc := func(a int, b int) int {
		return a + b
	}

	subFunc := func(a int, b int) int {
		return a - b
	}

	a := 20
	b := 35
	commonFunc, err := genCalculate(a, b, addFunc)
	if err != nil {
		fmt.Printf("err happend, err=%v\n", err)
		return
	}

	pc := reflect.ValueOf(addFunc).Pointer()
	f := runtime.FuncForPC(pc)

	fmt.Printf("%v(%v, %v)=%v\n", f.Name(), a, b, commonFunc(a, b))

	commonFunc, err = genCalculate(a, b, subFunc)
	if err != nil {
		fmt.Printf("err happend, err=%v\n", err)
		return
	}
	fmt.Printf("commonFunc(%v, %v)=%v\n", a, b, commonFunc(a, b))
}
