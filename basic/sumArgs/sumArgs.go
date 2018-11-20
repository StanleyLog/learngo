package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func sumArgs(numbers ... int) int {

	returnval := 0;
	for i := range numbers {
		returnval += i
	}
	return returnval
}

func main() {

	fmt.Println(sumArgs(1,2,3,4,5,6))
	p := reflect.ValueOf(sumArgs).Pointer()
	fmt.Println(runtime.FuncForPC(p).Name())
}
