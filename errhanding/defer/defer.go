package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)


func trydefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred.")
	return
	//defer fmt.Println(4)
}

func writeFile(filename string) {

	file, err := os.Create(filename)
	defer file.Close()

	if nil != err {
		if err, ok := err.(*os.PathError); ok {
			fmt.Printf("%s, %s, %s\n", err.Op, err.Path, err.Err)
		} else {
			panic(err)
		}
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	//trydefer()
	writeFile("/tmpa/test.log")
}
