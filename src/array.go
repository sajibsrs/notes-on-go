package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 100
	fmt.Println("set:", a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	fmt.Println("emp:", s)
	fmt.Println("emp:", len(s))
	fmt.Println("emp:", cap(s))
}
