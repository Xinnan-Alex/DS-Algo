package main

import (
	"fmt"
	"strings"
)

func mySolution(s string) int {
	arr := strings.Split(strings.TrimSpace(s), " ")
	lastIndex := arr[len(arr)-1]
	return len(lastIndex)
}

func main() {
	fmt.Println(mySolution("Hello World"))
}
