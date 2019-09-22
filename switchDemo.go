package main

import "fmt"

// switch表达式中，case子句的类型必须和switch中一致，同时case中的字面量不能重复
func main() {
	value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value3[4] {
	case 0, 1, 2:
		fmt.Println("0 or 1 or 2")
	case 2, 3, 4:
		fmt.Println("2 or 3 or 4")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}

}
