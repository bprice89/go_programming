package main

import "fmt"

func main() {
	//
	// for init; condition; post {
	//init = storing value, condition is ome condition, post is omething that happens after
	// }
	//initializing i as a value, loop keeps goin as
	//long as condiotion is met (i less than 10), after each time
	//through loop add 1 to i (i++)
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }
	//print inner loop inside outer loop
	// for i := 0; i <= 10; i++ {
	// 	for g := 6; g >= 3; g-- {
	// 		fmt.Println(g)
	// 	}
	// 	fmt.Println(i)
	// }
	//another way to print inner loop in outer loop
	for i := 0; i <= 10; i++ {
		for g := 6; g >= 3; g-- {
			fmt.Printf("The outer loop: %d\t The inner loop: %d\n", i, g)
		}

	}
	//another way to print inner loop in outer loop
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for g := 6; g >= 3; g-- {
			fmt.Printf("\t\t The inner loop: %d\n", g)
		}

	}

}
