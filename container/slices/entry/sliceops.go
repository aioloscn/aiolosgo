package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {

	var s []int
	for i := 0; i < 10; i++ {
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	printSlice(s)
	s1 := make([]int, 16)
	printSlice(s1)
	s2 := make([]int, 10, 32)
	printSlice(s2)
	copy(s2, s)
	printSlice(s2)
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	front := s2[0]
	s2 = s2[1:]
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(front, tail)
	printSlice(s2)
}
