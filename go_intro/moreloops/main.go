package main

import "fmt"

func main() {
	//for statement or while loop
	x := 1
	for x <= 10 {
		fmt.Println(x)
		x++
	}
	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println(y)
		y++
	}

	for a := 0; a < 20; a++ {
		if a%2 == 0 {
			fmt.Println(a, "even")
		}
		if a%2 != 0 {
			continue
		}
	}

}
