package main

import (
	"fmt"
)

type orange int
var x orange

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42

	fmt.Println(x)
}