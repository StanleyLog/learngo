package main

import (
	"fmt"
)

func main() {

	fmt.Println(00000)
	ch := make(chan int)

	//这里的协程必须先于ch <- 1, 因为如何进行协程的读写操作的话，那么就会立刻堵塞。这样的话下面的代码就不会被执行。
	go func(cha <-chan int) {

		// ch2 := make(chan int)
		<-cha
		// ch2 <- 1
		// return ch2
	}(ch)

	ch <- 1
	fmt.Println(11111)
}
