package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	channels := make([]chan int, 10)
// 	fmt.Println(channels)

// 	for i := 0; i < 10; i++ {
// 		channels[i] = make(chan int)
// 	}
// 	fmt.Println(channels)

// 	//send messager
// 	for i := 0; i < 10; i++ {
// 		go func(ch chan int, idx int) {
// 			ch <- idx
// 			fmt.Printf("Channel[%d] %v Send messager: %d\n", idx, ch,  idx)
// 		}(channels[i], i)
// 	}

// 	//receive messager
// 	for i := 0; i < 10; i++ {
// 		go func(ch chan int, idx int) {
// 			//fmt.Println(<-ch)
// 			fmt.Printf("Channel[%d] %v Received messager: %d, \n", idx, ch, <-ch)
// 		}(channels[i], i)
// 	}

// 	//for i := 0; i < 10; i++ {
// 	//
// 	//	channels[i] = make(chan int)
// 	//	channels[i] <- i
// 	//
// 	//	fmt.Println(i)
// 	//	//fmt.Printf("Send Channel[%d], Content: %d", i, i)
// 	//	//
// 	//	//go func(ii int, ch chan int) {
// 	//	//
// 	//	//	 flag := <- ch
// 	//	//	fmt.Printf("Reviced Channel[%d], Content: %d", ii, flag)
// 	//	//}(i, channels[i])
// 	//}

// 	time.Sleep(time.Minute)

// }
