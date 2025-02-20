package main

import "fmt"

func longestPalindrome(s string) string {
	res := ""
	for index := range s {
		left, right := index, index
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if len(s[left:right+1]) > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}

		left, right = index, index+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if len(s[left:right+1]) > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}
	}
	return res
}

func main() {
	s := "aacabdkacaa"

	fmt.Println(longestPalindrome(s))
}
