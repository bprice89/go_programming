package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("prints") //will stop here becasue it has not fallthrough
	case (4 == 4):
		fmt.Println(" should print")
	}

	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println(" should print")
	}

	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println(" should print")
		fallthrough
	case (15 == 365):
		fmt.Println("not true but prints becasue of fallthrouh")
		fallthrough
	default:
		fmt.Println("if nothing is true it will go to default")
	}

	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	default:
		fmt.Println("this is default")
	}

	switch "Bond" {
	case "James":
		fmt.Println("this should not print")
	case "rick":
		fmt.Println("this should not print")
	case "Bond":
		fmt.Println("this is james bond")
	}
}
