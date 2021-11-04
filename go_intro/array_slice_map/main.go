package main

import "fmt"

func main() {
	//array have to specify how many variables are in array, use slices
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)

	//slice
	//x := []type{values} //composite literal
	y := []int{4, 5, 6, 4, 2}
	fmt.Println(y)

	//range in slice
	b := []int{4, 5, 6, 354, 645, 23, 76, 1234, 634, 23}
	fmt.Println(b)
	fmt.Println(len(b))
	v := b[4]
	fmt.Println(v)
	//using range i = index in slice, v = value at that index, range b = loop through all of b
	for i, v := range b {
		fmt.Println(i, v)
		fmt.Printf("index is %v and value is %v\n", i, v)
	}

	//slicing a slice
	fmt.Println(b[2:6])
	fmt.Println(b[:8])
	fmt.Println(b[6:])

	//loop through all of slice without using range
	for i := 0; i < len(b); i++ {
		fmt.Println(i, b[i])

	}

	//appending slices
	z := []int{1, 2, 3, 6, 8, 9, 5}
	fmt.Println(z)
	z = append(z, 77, 88, 99, 101)
	fmt.Println(z)
	//append a slice to a slice
	m := []int{466, 588, 655, 89}
	fmt.Println(m)
	z = append(z, m...)
	fmt.Println(z)

	//deleting a slice
	z = append(z[:2], z[4:]...)
	fmt.Println(z)

	//slice make, you can make an array and have it save spots for later
	//10 is  current length of array, 100 is capacity of slice, if you append over capacity slice will double in size
	n := make([]int, 10, 100)
	fmt.Println(n)
	n[0] = 3
	n = append(n, 3423)
	fmt.Println(n)

	//multi dimensional slice
	multi := []string{"James", "bond", "frank", "george"}
	fork := []string{"ball", "blue", "red", "green"}

	multifork := [][]string{multi, fork}
	fmt.Println(multifork)

	//maps  map(string = key type)(int = value type)
	r := map[string]int{
		"james":     32,
		"bond":      45,
		"wolverine": 54,
	}

	fmt.Println(r)
	fmt.Println(r["bond"])

	//checking if key is in map
	v, ok := r["alex"]
	fmt.Println(v)
	fmt.Println(ok)

	//if statement to determine if key/value is in map
	if v, ok := r["james"]; ok {
		fmt.Println("if it was in the map it would print", v)
	}

	//add to map and use range

	r["todd"] = 34 //added to map

	for k, v := range r { //ranged through map
		fmt.Println(k, v)
	}

	//delete map entry
	delete(r, "todd")
	fmt.Println(r)

}
