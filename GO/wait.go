package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func processData(wg *sync.WaitGroup, result *[]int, data int) {
	lock.Lock()
	defer wg.Done()
	processingData := data * 2
	*result = append(*result, processingData)
	lock.Unlock()
}
func Test() {
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := []int{}

	for _, data := range input {
		wg.Add(1)
		go processData(&wg, &result, data)
	}

	wg.Wait()

	fmt.Println(result)
}
