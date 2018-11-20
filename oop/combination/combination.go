package main

import "fmt"

type A struct {
	Name string
}

func (a *A) SetName(name string)  {
	a.Name = name
}

type B struct {
	A
}

func (b *B) GetName() string {
	return b.A.Name
}

func main() {
	a := A{"Stanley"}
	b := B{ a}

	b.SetName("sdfsdfsdf")
	fmt.Println(b.GetName())
}