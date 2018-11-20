package main

import (
	"fmt"
)

func main() {

	var arr1 [5]int
	arr2 := [3]int {1,2,3}
	arr3 := [...]int {3,4,5}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v )
	}

	fmt.Println("----------------------")

	printArray(&arr2)
	printArray(&arr3)

	fmt.Println("======================")
	fmt.Print(arr2[0])
}

func printArray(arr *[3]int) {

	arr[0] = 100
	for _, v  := range arr {
		fmt.Println(v)
	}
	fmt.Println()
}
