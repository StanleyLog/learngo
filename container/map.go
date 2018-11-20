package main

import "fmt"

func main() {
	m1 := map[string]string {
		"first_name" : "Stanley",
		"last_name" : "Sun",
		"sex" : "F",
	}

	var m2 map[string]int

	m3 := make(map[string]string)

	fmt.Println(m1, m2, m3)


	fmt.Println("Traversing map")
	for k, v := range m1{
		fmt.Printf("key: %s, value: %s\n", k, v)
	}

	fmt.Println("Getting map")

	fmt.Println(m1["first_name"])
	fmt.Println(m1["last_name"])
	fmt.Println(m1["sex"])


	if value, isExist := m1["ssss"]; isExist {
		fmt.Println(value)
	} else {
		fmt.Printf("This value %s is not exist in map. ", "ssss")
	}

	fmt.Println("delete map ")

	delete(m1, "sex")
	fmt.Println(m1)

	fmt.Println(len(m1))

}
