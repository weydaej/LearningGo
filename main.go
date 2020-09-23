package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	var x int = 3 // can also be written x := 3
	x = 5
	y := 1
	sum := x + y
	fmt.Println(sum)

	var a [5]int
	a[0] = 99
	fmt.Println(a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	c := []int{1, 2, 3, 4, 5}
	c = append(c, 6)
	fmt.Println(c)

	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["dodecagon"] = 12

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])
	delete(vertices, "triangle")
	fmt.Println(vertices)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
