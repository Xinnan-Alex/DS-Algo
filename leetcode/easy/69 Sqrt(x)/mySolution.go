package main

import "fmt"

func mySqrt(x int) int {
	left, right := 0, x
	var result int

	for left <= right {
		midpoint := left + ((right - left) / 2)
		if midpoint*midpoint < x {
			left = midpoint + 1
			result = midpoint
		} else if midpoint*midpoint > x {
			right = midpoint - 1
		} else {
			return midpoint
		}
	}
	return result
}

func main() {
	fmt.Println(mySqrt(8))
}
