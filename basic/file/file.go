package main

import (
	"fmt"
	"io/ioutil"
)

const file_name = "/Users/stanley/hello.go"

func main() {

	contents, err := ioutil.ReadFile(file_name)
	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	//语法变形
	if contents, err := ioutil.ReadFile(file_name); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

}
