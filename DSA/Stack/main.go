package main

import "fmt"

type Stack struct {
	top int
	st  [5]int
}

func NewStack() *Stack {
	return &Stack{
		top: -1,
	}
}

func (s *Stack) push(val int) {
	if s.top == 4 {
		fmt.Println("Stack is full....")
		return
	}

	s.top++
	s.st[s.top] = val
}
func (s *Stack) getTop() (int, bool) {
	if s.top == -1 {
		fmt.Println("Satck is empty....")
		return -1, false
	}

	return s.st[s.top], true
}

func (s *Stack) pop() {
	if s.top == -1 {
		fmt.Println("Satck is empty")
		return
	}
	s.top -= 1
}

func (s *Stack) Print() {
	for i := 0; i <= s.top; i++ {
		fmt.Print(s.st[i], " ")
	}
	fmt.Println()
}

func (s *Stack) size() int {
	return s.top + 1
}

func main() {
	stack := NewStack()
	stack.push(10)
	stack.push(20)
	stack.push(30)
	stack.Print()
	stack.pop()

	top, ok := stack.getTop()
	if ok {
		fmt.Println("Top:", top)
	}
	fmt.Println("Size:", stack.size())

	stack.pop()

	top, ok = stack.getTop()
	if ok {
		fmt.Println("Top after pop:", top)
	}

	stack.Print()
}
