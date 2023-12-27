package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) // add 1 to the waitgroup
		go goSomething(i, &wg)
	}

	wg.Wait() // wait until the waitgroup counter is 0
}

func goSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done() // remove 1 from the waitgroup
	fmt.Printf("Started %d\n", i)
	time.Sleep(time.Second * 2)
	fmt.Println("Done!")
}
