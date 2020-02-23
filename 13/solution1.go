package main

import (
	"fmt"
	"strconv"
	"strings"
)

func romanToInt(s string) int {
	var res int
	temp := []int{}
	//映射
	Roman := make(map[string]int)
	Roman["I"] = 1
	Roman["V"] = 5
	Roman["X"] = 10
	Roman["L"] = 50
	Roman["C"] = 100
	Roman["D"] = 500
	Roman["M"] = 1000
	fmt.Printf("map is %v", Roman)
	//如果左边的字符表示的数字小于右边的，结果为右边的减去左边的，如果大于，则按位相加
	sep := ""
	arr := strings.Split(s, sep)
	for i := 0; i < len(arr)-1; i++ {
		if Roman[arr[i]] < Roman[arr[i+1]] {
			arr[i] = strconv.FormatInt(int64(Roman[arr[i+1]]-Roman[arr[i]]), 10)
			//arr[i+1] = ""
			//fmt.Printf("arr[i] is %v", arr[i])
			temp = append(temp, Roman[arr[i]])
			i = i + 2
			fmt.Printf("temp is %v", res)
		} else {
			temp = append(temp, Roman[arr[i]])
			//res = res + Roman[arr[i]]
		}
		//temp = append(temp, Roman[arr[i]])
		for i := 0; i < len(temp); i++ {
			res += temp[i]
		}
	}
	fmt.Printf("res is %v", res)
	return res
}
func main() {
	str := "MCMXCIV"
	r := romanToInt(str)
	fmt.Printf("res is %v", r)
}
