package main

import "fmt"

func main() {
	s := make(chan string, 2)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("sender hello", i)
			s <- fmt.Sprintf("receiver hello %d", i)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-s
		fmt.Println(val)
	}
}
