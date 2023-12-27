package main

import (
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)

	//fmt.Println(<-c1) // se queda bloqueo hasta que reciba un valor y luego imprime el valor c2
	//fmt.Println(<-c2)

	// Multiplexing
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			println("received", msg1)

		case msg2 := <-c2:
			println("received", msg2)
		}
	}
}

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
