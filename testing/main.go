package main

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return Fibonacci(x-1) + Fibonacci(x-2)
}
