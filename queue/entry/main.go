package main

import (
	"aiolosgo/queue"
	"fmt"
)

func main() {

	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	q.Pop()
	fmt.Println(q.IsEmpty())
	q.Push("aiolos")
	fmt.Println(q)
}
