package main

import "fmt"

// any global variables need to be used with var, walrus will not work outside of function
var y = 43

func main() {
	fmt.Println(y)
	x := 56
	fmt.Println(x)

	sub(400, 455)
}

func sub(a int, b int) {
	fmt.Println(a - b)
}
