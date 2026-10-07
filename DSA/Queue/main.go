package main

import "fmt"

type Que struct {
	currentSize int
	start       int
	end         int
	queue       []int
	capacity    int
}

func NewQueue(size int) *Que {
	return &Que{
		currentSize: 0,
		start:       -1,
		end:         -1,
		queue:       make([]int, size),
		capacity:    size,
	}
}
func (q *Que) push(data int) {
	if q.currentSize == q.capacity {
		fmt.Println("Queue is full....")
		return
	}

	if q.currentSize == 0 {
		q.start = 0
		q.end = 0
		q.queue[q.end] = data
		q.currentSize++
		fmt.Printf("%d push into queue , current capacity was %d, current size %d\n",
			data, q.capacity, q.currentSize)
	} else {
		q.end = (q.end + 1) % q.capacity
		q.queue[q.end] = data
		q.currentSize++
		fmt.Printf("%d push into queue , current capacity was %d, current size %d\n",
			data, q.capacity, q.currentSize)
	}
}

func (q *Que) pop() {
	if q.currentSize == 0 {
		fmt.Println("Queue is empty...nothing can be popped...")
		return
	}
	// 2 3 4 6
	// if 2 is pop, then new first ele was 3
	fmt.Println(q.queue[q.start], " removed")
	q.start = (q.start + 1) % q.capacity
	q.currentSize--
	if q.currentSize == 0 {
		q.start = -1
		q.end = -1
	}

}
func main() {
	queue := NewQueue(4)
	queue.push(1)
	queue.push(3)
	queue.push(4)
	queue.push(7)
	queue.pop()
	queue.pop()

	queue.push(10)

}
