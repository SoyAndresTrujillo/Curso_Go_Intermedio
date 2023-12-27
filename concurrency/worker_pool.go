package main

import "fmt"

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	numWorkers := 3

	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < numWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for i := 0; i < len(tasks); i++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker", id, "started job", job)
		fib := Fibonacci(job)
		fmt.Println("Worker", id, "finished job", job, "with result", fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
