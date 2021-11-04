package main

import "fmt"

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range s {
		fmt.Printf("index is %v and value is %v\n", i, v)
	}

	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:7])
	fmt.Println(s[1:6])

	s = append(s, 52)
	fmt.Println(s)
	s = append(s, 53, 54, 55)
	fmt.Println(s)

	y := []int{56, 57, 58, 59, 60}
	s = append(s, y...)
	fmt.Println(s)

	//deleting parts of a slice
	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	v := append(b[:3], b[6:]...)
	fmt.Println(v)

	d := make([]string, 50, 60)
	d = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	for i := 0; i < len(d); i++ {
		fmt.Println(i, d[i])
	}

	q := []string{"James", "Bond", "Shaken, not stirred"}
	w := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	qw := [][]string{q, w}
	fmt.Println(qw)
	for i, v := range qw {
		fmt.Println(i, v)
		for j, v2 := range v {
			fmt.Println(j, v2)
		}
	}

	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	//add map item
	m["oddJob"] = []string{"smaill thing", "killing", "hats"}

	for k, v := range m {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println("\t", k2, v2)
		}
	}
	//delete map entry
	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println("\t", k2, v2)
		}
	}
}
