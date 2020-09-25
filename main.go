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

	for key, value := range vertices {
		fmt.Println("key:", key, "value:", value)
	}

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])
	delete(vertices, "triangle")
	fmt.Println(vertices)

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// there are no while loops in Go o.O
	i := 0
	for i < 5 {
		fmt.Println("i:", i)
		i++
	}

	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	answer, even := numSum(2, 7)
	fmt.Println(answer, even)
}

func numSum(a int, b int) (int, bool) {
	isEven := false
	if (a+b)%2 == 0 {
		isEven = true
	}
	return a + b, isEven
}
