package main

func main() {
	//number := 10
	//fmt.Println(number)
	//fmt.Println(&number)

	// 1. Go routine
	//c1 := make(chan int)
	//go doSomething(c1)
	//<-c1

	// 2. Go routine
	//c2 := make(chan int)
	//go doSomethingSum(c1)
	//<-c1

	// 3. Go routine
	//c3 := make(chan int)
	//go doSomethingMultiplication(c1)
	//<-c1
}

//	func doSomething(c chan int) {
//		fmt.Println("Go routine doSomething")
//		c <- 1
//	}
//
//	func doSomethingSum(c chan int) {
//		fmt.Printf("Go routine doSomethingSum: %d \n", 1112312+113453678)
//		c <- 1
//	}
//func doSomethingMultiplication(c chan int) {
//	fmt.Printf("Go routine doSomethingMultiplication: %d \n", 12*1123123)
//	c <- 1
//}
