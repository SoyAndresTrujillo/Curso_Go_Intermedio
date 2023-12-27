package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Goroutine %d started\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n\n", i)
	<-c
}

//c := [][]
//c := [goRoutine1][goRoutine2]
//c := [goRoutine3][goRoutine2]
