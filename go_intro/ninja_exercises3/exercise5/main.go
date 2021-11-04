package main

import "fmt"

//Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("when %v is diveded by 4, the reaminder is %v\n", i, i%4)
	}
}
