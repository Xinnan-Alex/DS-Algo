package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	answer := 0
	// firstDigitIndex := -1
	sign := 1

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		sign = 1
		s = s[1:]
	}

	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			break
		}

		d := int(ch - '0')
		answer = answer*10 + d

		// math.MinInt32 < answer < math.MaxInt32
		if sign*answer < math.MinInt32 {
			return math.MinInt32
		} else if sign*answer > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return answer * sign
}

func main() {
	fmt.Println(myAtoi("9223372036854775808"))
}
