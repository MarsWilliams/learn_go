package main

import (
	"fmt"
)

type orange int

var x orange
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)

	x = 88
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
