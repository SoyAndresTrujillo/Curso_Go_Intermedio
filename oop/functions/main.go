package main

import (
	"fmt"
	"time"
)

func main() {
	//x := 1
	//y := func() int {
	//	return x * 2
	//}()
	//fmt.Println(y)
	//
	//z := 1
	//w := func() int {
	//	return z * 2
	//}()
	//fmt.Println(w)

	// At this moment, use the function anonymous broken the principle of DRY (Don't Repeat Yourself)

	// Function anonymous with goroutine's
	c := make(chan int)
	go func() {
		fmt.Println("Starting the goroutine")
		time.Sleep(5 * time.Second)
		fmt.Println("Finishing the goroutine")
		c <- 1
	}()
	<-c

}
