package strings

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "azAZ艾俄罗斯"
	fmt.Printf("len=%d\n", len(s))
	for i, ch := range []byte(s) {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println("\nRune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("%d%c ", i, ch)
	}
	fmt.Println()
}
