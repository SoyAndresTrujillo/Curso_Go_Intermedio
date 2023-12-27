package main

import "fmt"

func main() {
	// channel unbuffered
	//c := make(chan int)

	// channel buffer (3)
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
