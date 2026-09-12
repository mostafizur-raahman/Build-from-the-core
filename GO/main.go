package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int) {
	duration := time.Duration(100+rand.Intn(400)) * time.Millisecond
	time.Sleep(duration)
	fmt.Printf("worker %d slept for %v\n", id, duration)
}

func main() {
	ch := make(chan struct{})
	bh := make(chan struct{})
	go func() {
		fmt.Println("working....")
		time.Sleep(time.Second)
		fmt.Println("done!")
		ch <- struct{}{}
	}()

	go func() {
		fmt.Println("working1....")
		time.Sleep(time.Second)
		fmt.Println("done1!")
		bh <- struct{}{}
	}()

	res := <-ch // block untill received
	fmt.Println("now ch is unblock ", res)
	res2 := <-bh
	fmt.Println("now bh is unblock ", res2)
	fmt.Println("All !done")
}
