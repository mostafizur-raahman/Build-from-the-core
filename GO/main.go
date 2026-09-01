package main

import (
	"fmt"
	"time"
)

func makeAccumulator() func(n int) int {
	total := 0
	return func(n int) int {
		total += n
		return total
	}
}

func timed(name string, fn func()) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("%s took %v\n", name, elapsed)
	}()

	fn()
}

func timed2(name string, fn func()) {
	start := time.Now()

	fn()
	elapsed := time.Since(start)
	fmt.Printf("%s took %v\n", name, elapsed)

}

func main() {
	acc := makeAccumulator()
	fmt.Println(acc(5))
	fmt.Println(acc(10))
	fmt.Println(acc(3))

	// timed("Test 1: ", func() {
	// 	time.Sleep(100 * time.Millisecond)
	// })
	timed("panicking operation", func() {
		time.Sleep(50 * time.Millisecond)
		panic("oh no!")
	})

}
