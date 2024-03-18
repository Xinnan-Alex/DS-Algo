package main

func plusOne(digits []int) []int {
    for i:=len(digits): i<=0; i--{
		fmt.println(digits[i])
	}
}

func main(){
	plusOne([1,2,3])
}