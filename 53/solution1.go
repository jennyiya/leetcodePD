package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	dp:=nums[0]
	max := nums[0]
	for i:=1;i<len(nums);i++{
        if dp > -1{
			dp = dp +nums[i]
            }else{
                dp = nums[i]
            }
            max = int(math.Max(float64(dp), float64(max)))
	}   
	return max
}
func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	c := maxSubArray(a)
	fmt.Printf("the res is %v", c)
}
