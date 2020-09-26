package main1

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Emily", age: 23}
	fmt.Println(p)
}
