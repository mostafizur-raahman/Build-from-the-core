package main

import "fmt"

type Queue struct {
	currentSize int
	queue       []int
	start       int
	end         int
	sz          int
}

func NewQueue(size int) *Queue {
	return &Queue{
		currentSize: 0,
		queue:       make([]int, size),
		start:       -1,
		end:         -1,
		sz:          size,
	}
}

func (q *Queue) push(data int) {
	if q.currentSize == q.sz {
		fmt.Println("Queue is full...Can't push until some pop...")
		return
	}

	if q.currentSize == 0 {
		q.start = 0
		q.end = 0
		q.queue[q.start] = data
		q.currentSize++
		fmt.Println(data, "Pushed into Queue, position at ", q.currentSize)
	} else {
		q.end = (q.end + 1) % q.sz
		q.queue[q.end] = data
		q.end++
		q.currentSize++
		fmt.Println(data, "Pushed into Queue, position at ", q.currentSize)
	}
}
func (q *Queue) pop() {
	if q.currentSize == 0 {
		fmt.Println("Queue is empty you can not pop, right now , push befor pop")
		return
	} else {
		if q.currentSize == 1 {
			q.currentSize--
			q.start = -1
			q.end = -1
		} else {
			q.start = (q.start + 1) % q.sz
		}
	}
}

func (q *Queue) front() (int, bool) {
	if q.currentSize == 0 {
		return -1, false
	}
	return q.queue[q.start], true
}

func (q *Queue) isEmpty() bool {
	return q.currentSize == 0
}
func main() {
	queue := NewQueue(2)
	queue.push(2)
	queue.push(3)
	queue.pop()
	queue.push(5)
	res, ok := queue.front()
	if ok {
		fmt.Println("Front is ", res)
	}
}
