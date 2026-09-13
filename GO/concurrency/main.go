package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func process(data int) int {
	time.Sleep(time.Second * 3)
	return 2 * data
}
func processData(wg *sync.WaitGroup, index *int, data int) {
	defer wg.Done()
	processData := process(data)
	*index = processData
}
func Concurrent() {
	start := time.Now()

	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	output := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)
		go processData(&wg, &output[i], data)
	}

	wg.Wait()
	end := time.Since(start)
	fmt.Println("Result is , ", output)
	fmt.Println("Time taken : ", end)

}
