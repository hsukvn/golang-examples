package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	foo := add
	val1 := add(1, 2)
	val2 := foo(3, 4)
	fmt.Printf("%v\n", val1)
	fmt.Printf("%v\n", val2)
	fmt.Printf("%v\n", add(5, 6))
}
