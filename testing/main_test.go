package main

import "testing"

func TestSum(T *testing.T) {
	//result := Sum(52, 5)
	//if result != 10 {
	//	T.Errorf("Expected 10, got %d", result)
	//}

	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
	}

	for _, item := range tables {
		total := Sum(item.x, item.y)
		if total != item.n {
			T.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", item.x, item.y, total, item.n)
		}
	}
}

func TestGetMax(T *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{2, 1, 2},
		{3, 2, 3},
		{4, 6, 6},
	}

	for _, item := range tables {
		total := GetMax(item.x, item.y)
		if total != item.n {
			T.Errorf("GetMax of (%d+%d) was incorrect, got: %d, want: %d.", item.x, item.y, total, item.n)
		}
	}
}

func TestFibonacci(T *testing.T) {
	tables := []struct {
		x int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.x)
		if fib != item.n {
			T.Errorf("Fibonacci of (%d) was incorrect, got: %d, want: %d.", item.x, fib, item.n)
		}
	}
}
