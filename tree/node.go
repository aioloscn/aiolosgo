package tree

import "fmt"

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

func (node TreeNode) Print() {
	fmt.Println(node.Value)
}

func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("node is nil")
		return
	}
	node.Value = value
}

func createTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value} // 局部变量分配在栈上，函数退出后立刻销毁
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

