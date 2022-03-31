package main

import "fmt"

type Point struct {
	X, Y int
}

func setX(p *Point) *Point {
	p.X += 3
	return p
}

func getPosition(x, y int) Point {
	return Point{x, y}
}

func main() {
	position := &Point{2, 3}
	p := setX(position)
	fmt.Println(&position)
	fmt.Println(&position.X)
	fmt.Println(&position.Y)
	fmt.Println(p)
}
