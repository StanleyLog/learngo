package main

import (
	"fmt"
	"learngo/oop/inter"
)

func main() {

	var t inter.Getter

	t = &inter.Text{"sdf"}
	//fmt.Println(t.Get())

	//fmt.Printf("%T %v\n", t, t)

	//a := t.(inter.Text)

	if ttt, ok := t.(inter.Getter); ok {

		fmt.Println("---------")
		fmt.Println(ttt)
		fmt.Println(ok)
	}else {

		fmt.Println("=========")
		fmt.Println(ttt)
		fmt.Println(ok)
	}
	//
	//
	//
	//switch tt := t.(type) {
	//
	//case inter.Getter:
	//	fmt.Println("inter.getter")
	//case *inter.Text:
	//	fmt.Println("inter.textgeter")
	//default:
	//	fmt.Println(tt)
	//}
	//
	//a := t.(*inter.Text)
	//fmt.Println(a)
	//
	//
	//b := []interface{}{1,2,3,4,5}
	//fmt.Println(b[0].(int))

}
