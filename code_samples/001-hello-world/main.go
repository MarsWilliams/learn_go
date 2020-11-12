package main

import "fmt"

func main() {

	foo()

	fmt.Println("Hello World")

	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
		fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("And then we exited")
}
