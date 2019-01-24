package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Fibonacci() intGen {
	a, b := 0, 1					// 0
	return func() int {
		a, b = b, a + b					// 3
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()								// 2
	if (next > 10000) {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)			// 4
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)				// 1
	for scanner.Scan() {
		fmt.Println(scanner.Text())					// 5
	}
}

func main() {
	f := Fibonacci()
	printFileContents(f)
}
