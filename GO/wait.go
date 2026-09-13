package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func process(data int) int {
	time.Sleep(2 * time.Second)
	return 2 * data
}
func processData(wg *sync.WaitGroup, result *[]int, data int) {
	defer wg.Done()
	processingData := process(data)

	lock.Lock()
	*result = append(*result, processingData)
	lock.Unlock()
}
func Test() {
	start := time.Now()

	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := []int{}

	for _, data := range input {
		wg.Add(1)
		go processData(&wg, &result, data)
	}

	wg.Wait()
	end := time.Since(start)
	fmt.Println(result, end)
}
