package fib

func Fibonacci() func() int {
	a, b := 0, 1					// 0
	return func() int {
		a, b = b, a + b					// 3
		return a
	}
}
