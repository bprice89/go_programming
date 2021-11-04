package main

import "fmt"

//print every rune code point of the uppercase alphaabet 3 times
func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
	}
}
