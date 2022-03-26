package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	s = append(s, "d")

	fmt.Println("s:", s)
	fmt.Println("s len:", len(s))
	fmt.Println("s cap:", cap(s))

	s1 := []int{}
	s1 = append(s1, 1)

	fmt.Println("s1:", s1)
	fmt.Println("s1 len:", len(s1))
	fmt.Println("s1 cap:", cap(s1))
}
