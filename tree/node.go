package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("node is nil")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value} // 局部变量分配在栈上，函数退出后立刻销毁
}

func (node *Node) Traversal() {
	node.TraversalFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraversalFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraversalFunc(f)
	f(node)
	node.Right.TraversalFunc(f)
}

