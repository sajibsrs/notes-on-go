package main

import "fmt"

func main() {
	arr_one := [...]int{1, 2, 3}
	var arr_two [3]string

	arr_two[0] = "a"
	arr_two[1] = "b"
	arr_two[2] = "c"

	fmt.Println("arr_one:", arr_one)
	fmt.Println("arr_two:", arr_two)
}
