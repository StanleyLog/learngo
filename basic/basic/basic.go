package main

import "fmt"

var (
	a = 4
	b = 2
)

func variable(){
	var  a int
	var  b string
	fmt.Printf("%d %q\n", a, b)
}

func variableTypeDeduction(){
	var a , b, c , d, e = 1, '2', true, "sdf", "a"
	fmt.Println(a, b, c, d, e)
}
func variableShorter(){
	a, b, c := 1, '2', "sdf"
	fmt.Println(a, b, c)
}

func main() {
	//fmt.Println("sdfsdflskdjflj")
	variableTypeDeduction()
	variableShorter()


	var a float32
	a = float32(float32(1)/3)

	fmt.Println(a)



	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b , kb, mb, gb, tb, pb)

}