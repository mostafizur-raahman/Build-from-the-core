package main

import (
	"fmt"
	"time"
)

func generate1(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, n := range nums {
			fmt.Println("GENERATE:", n)

			time.Sleep(1 * time.Second)
			out <- n
		}
	}()

	return out
}

func square1(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {

			time.Sleep(1 * time.Second)

			result := n * n
			fmt.Println("SQUARE:", result)

			out <- result
		}
	}()

	return out
}

func pileline() {
	nums := generate(1, 2, 3, 4)
	squares := square(nums)

	for n := range squares {
		fmt.Println("MAIN: printing", n)
	}
}
