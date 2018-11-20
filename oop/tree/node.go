package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

func (node *Node) SetValue(value int) {
	node.Value = value

}

func (node *Node) SetLeftNode(leftNode *Node)  {
	node.Left = leftNode
}

func (node *Node) SetRightNode(rightNode *Node)  {
	node.Right = rightNode
}

func (node *Node) Traverse() {

	if nil == node {
		return
	}

	node.Left.Traverse()
	fmt.Println(node.Value)
	node.Right.Traverse()
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
