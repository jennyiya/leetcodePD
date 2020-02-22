package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	var b int
	if len(digits) == 0 {
		new := []int{1}
		return new
	}
	if len(digits) == 1 {
		if digits[0] != 9 {
			digits[0] = digits[0] + 1

		} else {
			digits := []int{1, 0}
			return digits
		}
		return digits
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			b++
		}
		fmt.Printf("b is %v", b)
	}
	if b == 0 {
		new := []int{1}
		return new
	} else {
		if b == len(digits) {
			var temp []int
			
			for i := 1; i < len(digits)+1; i++ {
				temp[i] = 0
				temp[0] = 1
			}
			fmt.Printf("the res is %v", temp)
			return temp
		} else {
			i := len(digits) - 1
			if digits[len(digits)-1] == 9 {
				digits[i] = 0
				digits[i-1] = digits[i-1] + 1
			} else {
				digits[i] = digits[i] + 1
			}
		}
	}
	return digits
}

func main() {
	a := []int{9, 9, 9}
	res := plusOne(a)
	fmt.Printf("res is %v\n", res)
}
