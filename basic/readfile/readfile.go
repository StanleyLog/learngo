package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/stanley/hello.go")

	if nil != err {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	file_num := 0
	for scanner.Scan() {
		file_num++
		fmt.Println(file_num, scanner.Text())
	}
}
