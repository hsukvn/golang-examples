package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	j := 0
	for j < 0 {
		j += 1
	}
	fmt.Println(j)
}
