package main

import (
	"fmt"
	"reflect"
)

func main() {
	// printf %T flag
	var x interface{} = 23.3
	fmt.Printf("%T\n", x)

	// type switch
	var x2 interface{} = 23
	switch v := x2.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	case bool:
		fmt.Println(v)
	case chan int:
		fmt.Println(v)
	default:
		fmt.Println("unknown")
	}

	// reflect package
	var x3 interface{} = make(chan int)
	fmt.Println(reflect.TypeOf(x3))

}
