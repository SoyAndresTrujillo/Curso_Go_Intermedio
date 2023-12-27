package main

import "fmt"

// Generator have generator only writing channel
func Generator(generator chan<- int) {
	for i := 1; i <= 10; i++ {
		generator <- i
	}
	close(generator)
}

// Double have generator reading channel and double writing channel
func Double(generator <-chan int, double chan<- int) {
	for value := range generator {
		double <- 2 * value
		//generator <- 1
	}
	close(double)
}

func Print(generator chan int) {
	for value := range generator {
		fmt.Println(value)
	}
}

func main() {
	generator := make(chan int)
	double := make(chan int)

	go Generator(generator)
	go Double(generator, double)

	Print(double)
}
