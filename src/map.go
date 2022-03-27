package main

import "fmt"

func main() {
	mp := map[string]string{}
	mp["hello"] = "world"
	mp["good"] = "morning"

	hlw := mp["hello"]

	fmt.Println("mp: ", mp)
	fmt.Println("mp: ", len(mp))
	fmt.Println("hlw: ", hlw)

	mp1 := make(map[string]string)

	mp1["hello"] = "mom"
	mp1["hello"] = "dad"

	fmt.Println("mp1: ", mp1)
	fmt.Println("mp1: ", len(mp1))
	fmt.Println("mp1: ", mp1["hello"])
}
