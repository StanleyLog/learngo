package main

import (
	"fmt"
	"math/rand"
	"sync"
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
	lock := sync.Mutex{}
	func() {
		go func() {
			i := 0
			for {
				lock.Lock()
				time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
				out <- i
				i++
				lock.Unlock()
			}
		}()
	}()
	return out
}

func main() {

	var c1, c2 = generator(), generator()
	w := createWrok(0)

	tm := time.After(10 * time.Second)
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
		case <-tm:
			fmt.Println("This program is shutdown.")
		}
	}

}
