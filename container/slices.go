package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 50
}

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	fmt.Println(s)
	updateSlice(s)
	fmt.Println(s)
	fmt.Println(arr)
	s2 := s[:5]
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	s2 = s2[2:]
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	s3 := append(s2, 11)
	s4 := append(s3, 20)
	s5 := append(s4, 33)
	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	fmt.Println(arr)
}
