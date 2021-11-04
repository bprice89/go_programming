package main

import (
	"fmt"
)

func main() {

	x := 42
	fmt.Println("hello")

	fmt.Println(x + x)
	add(100, 564)
	foo()
	s := "hello"
	fmt.Println(s)
	s = "goodbye"
	fmt.Println(s)

	// variables.sub(555, 4544)
	// for i := 0; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func foo() {
	fmt.Println("Im in foo")
}
