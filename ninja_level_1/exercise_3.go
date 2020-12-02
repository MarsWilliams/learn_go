package main

import (
	"fmt"
)

var x int = 24
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)

	fmt.Println(s)
}
