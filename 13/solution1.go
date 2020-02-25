package main

import (
	"fmt"
)

func romanToInt(s string) int {
	var sum int
	//temp := []int{}
	//映射
	Roman := make(map[byte]int)
	Roman['I'] = 1
	Roman['V'] = 5
	Roman['X'] = 10
	Roman['L'] = 50
	Roman['C'] = 100
	Roman['D'] = 500
	Roman['M'] = 1000
	//如果左边的字符表示的数字小于右边的，结果为右边的减去左边的，如果大于，则按位相加
	
	for i:=len(s)-1;i>=0;i--{
		arr := Roman[s[i]]
		if i>0 && Roman[s[i]] > Roman[s[i-1]]{
			arr -= Roman[s[i-1]]
			i--
		}
		sum += arr
	}
	return sum
}
func main() {
	str := "MCMXCIV"
	r := romanToInt(str)
	fmt.Printf("res is %v", r)
}
