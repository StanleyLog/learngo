package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	s := "This is a sample, 这是一个测试例子"


	//b := []byte(s)

	//fmt.Println(s)
	//fmt.Println(b)

	for i, ch := range s {
		fmt.Printf("index: %d, ch: %c(%X)\n", i, ch, ch)
	}

	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]

		fmt.Printf("%c", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("index: %d, ch: %c(%d) \n", i, ch, ch)
	}
	fmt.Println()


	s = " This is a sample2 "

	fmt.Println(strings.Trim(strings.ToUpper(s), " "))
	for i, v := range strings.Fields(s){
		fmt.Println(i, v)
	}
	fmt.Println()

	for i, v := range strings.Split(s, "i"){
		fmt.Println(i, v)
	}
	fmt.Println()

	arary_s := []string {
		"this",
		"is",
		"a",
		"sample3",
	}
	s = strings.Join(arary_s, "_")
	fmt.Println(s)

	if strings.Contains(s, "is") {
		index := strings.Index(s,"is")
		fmt.Printf("%s is contains the '%s', is at %d", s, "is", index)
	}
}
