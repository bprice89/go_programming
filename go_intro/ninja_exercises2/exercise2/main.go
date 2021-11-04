package main

import "fmt"

func main() {
	a := 42 == 42
	b := 43 <= 67
	c := 34 >= 543
	d := 56 != 56
	e := 67 < 89
	f := 89 > 93

	fmt.Println(a, b, c, d, e, f)
}
