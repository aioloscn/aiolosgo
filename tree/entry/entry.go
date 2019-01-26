package main

import (
	"aiolosgo/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {

	var root tree.Node
	root.Traversal()
	root = tree.Node{Value: 3}
	root.Left = new(tree.Node)
	root.Right = &tree.Node{}
	root.Left.SetValue(1)
	root.Right.SetValue(180)
	//
	//nodes := []tree.Node {
	//	{5, nil, nil},
	//	{},
	//	{10, nil, &root},
	//}
	//fmt.Println(nodes)
	//
	//cur := tree.CreateNode(99)
	//fmt.Println(cur)
	//root.Left.SetValue(12)
	//root.Left.Print()
	//pRoot := &root
	//pRoot.SetValue(102)
	//root.Print()
	//var test *tree.Node
	//test.SetValue(345)
	//
	//root.Traverse()
	//myRoot := myTreeNode{&root}
	//myRoot.postOrder()

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		fmt.Println("receive data")
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)
}
