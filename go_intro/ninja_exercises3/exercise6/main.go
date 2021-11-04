package main

import "fmt"

func main() {
	x := 76
	if x == 45 {
		fmt.Println("x equals 45")
	} else if x == 76 {
		fmt.Println("x equals 76")

	}
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("watersport")
	case "surfing":
		fmt.Println("surfiing")

	}
}
