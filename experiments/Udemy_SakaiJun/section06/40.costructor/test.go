package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func Area(v Vertex) int {
	return v.x * v.y

}

func New
func main() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))
}
