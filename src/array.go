package main

import "fmt"

func main() {
	arrOne := [...]int{1, 2, 3}
	var arrTwo [3]string

	arrTwo[0] = "a"
	arrTwo[1] = "b"
	arrTwo[2] = "c"

	fmt.Println("arr_one:", arrOne)
	fmt.Println("arr_two:", arrTwo)
}
