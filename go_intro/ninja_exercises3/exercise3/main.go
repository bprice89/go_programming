package main

import "fmt"

//use for condiotion to print all years I have been alive
func main() {
	x := 2021
	for x >= 1985 {
		fmt.Println(x)
		x--
	}
	//or
	y := 1985
	for y <= 2021 {
		fmt.Println(x)
		x++
	}

}
