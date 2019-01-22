package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 50
}

func main() {

	arr := [5]int{1, 2, 5, 7, 9}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	str := "aiolos"
	for i := range str {
		fmt.Printf("%q", str[i])
	}
	printArray(&arr)
	fmt.Println(arr)
}
