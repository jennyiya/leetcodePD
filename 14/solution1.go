package main

import (
	//"strings"
	"fmt"
)
func longestCommonPrefix(strs []string) string {
		if len(strs) == 0 {
			return ""
		} else if len(strs)==1{
			return strs[0]
		}else{
			minlen := len(strs[0])
			for i := 1; i < len(strs); i++ {
				if len(strs[i]) <= minlen {
					minlen = len(strs[i])
				}
			}
			//var k int
			res := ""
			for j := 0; j < minlen; j++ {
				flag := true
				val := strs[0][j]
				for _, v := range strs[1:] {
					if val != v[j] {
						flag = false
						break
					}
				}
				if flag == false && j != 0 {
					res = strs[0][:j]
					break
				}else if flag == false && j == 0 {
					res = ""
					break
				}else{
					res = strs[0][:j+1]
				}
	
			}
			return res
		}
	}
func main() {
	a := []string{"d","d"}
	c := longestCommonPrefix(a)
	fmt.Printf("the res is %v", c)
}
