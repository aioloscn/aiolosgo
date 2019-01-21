package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

func triangle() {
	const a, b = 3, 4
	c := int(math.Sqrt(a * a + b * b))
	fmt.Println(c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	euler()
	triangle()
	enums()
}
