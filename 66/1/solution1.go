package main

import (
	"fmt"
	"math"
)

func plusOne(digits []int) []int {
	var num int
	for i := len(digits) - 1; i >= 0; i-- {
		num = num + digits[i]*int(math.Pow10(len(digits)-i-1))
	}
	fmt.Printf("the num is%v", num)
	num = num + 1
	var slice []int
	for i := 0; i < len(digits); i++ {
		//p=(int(math.Pow10(len(digits)-i-1)))
		num = num - (num / (int(math.Pow10(len(digits) - i - 1))) * int(math.Pow10(len(digits)-i-1)))
		slice = append(slice, num/(int(math.Pow10(len(digits)-i-1))))

	}
	return slice
}
func main() {
	a := []int{4, 3, 2, 1}
	res := plusOne(a)
	fmt.Printf("the res is %v", res)
}
