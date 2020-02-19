package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var i, j, k int
	var l []int
	var temp [][]int
	for i = len(nums) - 1; i >= 2; i-- {
		for j = i - 1; j >= 1; j-- {
			for k = j - 1; k >= 0; k-- {
				fmt.Printf("res= %v ", nums[i]+nums[j]+nums[k])
				if nums[i]+nums[j]+nums[k] == 0 {
					l = []int{nums[i], nums[j], nums[k]}
					sort.Ints(l)
					temp = append(temp, l)
					for a := 0; a < len(temp); a++ {
						for b := a + 1; b < len(temp); b++ {
							i := 0
							for c := 0; c < len(temp[a]); c++ {
								if temp[a][c] != temp[b][c] {
									//fmt.Printf("no dup %v \n", b)
									break
								} else {
									i = i + 1
								}
								if i == len(temp[a]) {
									temp[a] = nil
									fmt.Printf("%v is dup", temp[b])
									break
								}
							}

						}
					}
				}
			}
		}
		fmt.Printf("j is %v ", j)
	}
	fmt.Printf("i is %v \n", i)
	fmt.Printf("the temp is %v", temp)
	return temp
}

//fmt.Printf("the temp is %v",temp)

func main() {

	var nums []int
	nums = []int{-1, 0, 1, 2, -1, -4}
	//threeSum(nums)
	res := [][]int{}
	res = threeSum(nums)
	fmt.Printf("the res %v", res)
}
