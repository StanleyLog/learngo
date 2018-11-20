package main

import (
	"fmt"
	"learngo/oop/tree"
)

//import "fmt"

func main() {

	var root tree.Node

	fmt.Println(root)

	root = tree.Node{}
	fmt.Println(root)

	root.Left = &tree.Node{Value:3}
	root.Right = &tree.Node{5, nil , nil }
	fmt.Println(root)

	root.Right.Left = new(tree.Node)
	fmt.Println(root)

	nodes := []tree.Node{
		{},
		{5,nil, nil},
		{Value:6},
	}
	fmt.Println(nodes)

	nodes = append(nodes, *tree.CreateNode(4))
	fmt.Println(nodes)

	root.SetValue(100)
	fmt.Println(root)

	root.Traverse()

}
