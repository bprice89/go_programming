package main

import "fmt"

func main() {
	if true {
		fmt.Println("I am True1")

	}
	if false {
		fmt.Println("I am false2")
	}
	if !true {
		fmt.Println("I am false3")
	}
	if !false {
		fmt.Println("i am true4")
	}

	x := 42
	if x == 40 {
		fmt.Println("x is not 40")
	} else if x == 41 {
		fmt.Println("x is not 41")
	} else {
		fmt.Println("keep guessing")
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, "divisible by 2")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "divisible by 5 and 3")
		} else if i%3 == 0 {
			fmt.Println(i, "divisible by 3")
		} else if i%5 == 0 {
			fmt.Println(i, "divisible by 5")
		}
	}

}
