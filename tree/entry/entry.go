package main

import "fmt"

func main() {

	var root tree.TreeNode
	root = tree.TreeNode{Value: 3}
	root.Left = new(TreeNode)
	root.Right = &TreeNode{}
	root.Left.SetValue(1)
	root.Right.SetValue(180)

	nodes := []TreeNode {
		{5, nil, nil},
		{},
		{10, nil, &root},
	}
	fmt.Println(nodes)

	cur := CreateTreeNode(99)
	fmt.Println(cur)
	root.Left.SetValue(12)
	root.Left.Print()
	pRoot := &root
	pRoot.SetValue(102)
	root.Print()
	var test *TreeNode
	test.SetValue(345)

	root.Traverse()
}
