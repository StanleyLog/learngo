package main

import (
	"fmt"
)

func main() {
	//
	//var a int = 128
	//fmt.Println(a)
	//
	//b := byte(a)
	//fmt.Println(b)
	//
	//c := rune(b)
	//fmt.Println(c)
	//
	//d := string(b)
	//fmt.Println(d)

	//fmt.Println(math.Pow(8, 4))

	str := "Test,测试。"
	//b := []byte(str)
	//
	//for i := 0; i < len(b); i++ {
	//	fmt.Printf("%d: %q \t %X[%s]\n", i, b[i], b[i], biu.ByteToBinaryString(b[i]))
	//}

	for i, v := range str {
		fmt.Printf("%d, %q [%  X]\n", i, v, []byte(string(v)))
	}
	//
	//r := []rune(str)
	//for j := 0; j < len(r); j++ {
	//	fmt.Println(j, r[j], string(r[j]))
	//}
	//
	//fmt.Println()
	//fmt.Println(string(97))


	//a := []byte(str)
	//fmt.Printf("The string: %q\n", str)
	//fmt.Printf("  => runes(char)[%s]: %q\n", biu.BinaryStringToBytes(a), []rune(str))
	//fmt.Printf("  => runes(hex): %x\n", []rune(str))
	//fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))



	//b := byte(0x1)
	//fmt.Printf("%d\n",b)
	//fmt.Println(biu.ByteToBinaryString(b))


	//for i := 0; i < int(math.Pow(2, 8*2)); i++ {
	//	fmt.Printf("Code= %d \t rune=%c \n", i, rune(i))
	//}


}
