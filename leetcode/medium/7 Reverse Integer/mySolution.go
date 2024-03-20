package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var answer int
	var negative bool

	// check if x is negative, keep note and convert to positive
	if x < 0 {
		negative = true
		x *= -1
	}

	for x > 0 {
		a := x % 10
		answer *= 10
		answer += a
		x /= 10
	}

	// convert back to negative
	if negative {
		answer *= -1
	}

	// DO subtraction between the reversed int and max of int32
	// if it's positive means its overflowing
	if (answer - math.MaxInt32) > 0 {
		return 0
	}

	return answer
}

func main() {
	fmt.Println(reverse(1534236469))
}
