package main

import "fmt"

//creating my own type (hotdog) and making that type an int
type hotdog int

var x hotdog

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
