package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s string
		ans int
	} {
		// Normal cases
		{"abcabcab", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"bbbb", 1},
		{"abcabcabcd", 4},
		{"是度始发地搜喝啊好的的艾尔拱趴奇偶", 10},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "是度始发地搜喝啊好的的艾尔拱趴奇偶"
	ans := 10

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
