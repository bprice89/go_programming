package main

import "fmt"

func main() {
	x := 45
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	//bit shifting x variable one spot to the left, declaring y a new variable
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}
