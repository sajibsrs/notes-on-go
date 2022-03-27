package main

import "fmt"

func main() {
	mp := map[string]string{}
	mp["hello"] = "world"
	mp["good"] = "morning"

	delete(mp, "good")

	hlw, st := mp["hello"]

	fmt.Println("mp: ", mp)
	fmt.Println("mp: ", len(mp))
	fmt.Println("hlw: ", hlw, st)

	mp1 := make(map[string]string)

	mp1["hello"] = "mom"
	mp1["hello"] = "dad"

	fmt.Println("mp1: ", mp1)
	fmt.Println("mp1: ", len(mp1))
	fmt.Println("mp1: ", mp1["hello"])

	mp2 := map[string]string{"nice": "weather", "lovely": "environment"}

	val, st1 := mp2["nice"]

	fmt.Println("val and status: ", val, st1)
}
