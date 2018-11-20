package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work(id int, c chan int) {
	for n := range c {
		fmt.Printf("Work %d received %d\n", id, n)
	}
}

func createWrok(id int) chan int {
	c := make(chan int)
	go work(id, c)
	return c
}

func generator() chan int {

	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {

	var c1, c2 = generator(), generator()
	w := createWrok(0)

	for {
		select {
		case n := <-c1:
			// fmt.Println("Received from c1: ", n)
			w <- n
		case n := <-c2:
			// fmt.Println("Received from c2: ", n)
			w <- n
			// default:
			// 	fmt.Println("No value received.")
		}
	}

	fmt.Println("sdfsdfsdf")
}
