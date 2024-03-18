package main

import "fmt"

func improvedSolution(digits []int) []int {
	// we start from the last index
	for i := len(digits) - 1; i >= 0; i-- {
		// if the current digit is smaller than 9, eg 1,2,3,4 5...8
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// else if its 9, 9+1 = 10 so change current digit to 0
		digits[i] = 0
	}

	// if digits[0] = 0 means the slice is out of bound eg [9,9,9] -> [0,0,0]
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(improvedSolution([]int{8, 9, 9, 9}))
}
