package main

import (
	"fmt"
	"learngo/oop/queue"
)

func main() {
	q := queue.Queue{0}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q)

	a := queue.Queue.IsEmpty(nil)
	fmt.Println(a)

}


