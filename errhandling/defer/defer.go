package main

import (
	"aiolosgo/functional/fib"
	"bufio"
	"fmt"
	"os"
)

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")

	for i := 0; i < 30; i++ {
		defer fmt.Println(i)
	}
}
