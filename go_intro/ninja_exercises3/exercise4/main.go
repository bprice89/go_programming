package main

import "fmt"

//print every year i was born using a for{} loop

func main() {
	x := 1984

	for {

		if x == 2022 {
			break
		}
		fmt.Println(x)
		x++
	}
}
