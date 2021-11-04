package main

import "fmt"

//untyped and typed const
const (
	a     = 42
	b int = 56
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
}
