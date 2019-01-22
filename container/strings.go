package main

import "fmt"

func main() {

	s := "azAZ艾俄罗斯"
	fmt.Println(s)
	for _, ch := range []byte(s) {
		fmt.Printf("%X ", ch)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
}
