package main

import "fmt"

func generate(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}
func main() {
	nums := generate(1, 2, 3, 4)
	squares := square(nums)

	for n := range squares {
		fmt.Println(n)
	}
}
