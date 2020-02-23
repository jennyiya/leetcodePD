package main

import (
	//"strings"
	"fmt"
)
func longestCommonPrefix(strs []string) string {
	var a int
	var pre string
	if len(strs) == 0 || len(strs) == 1 {
		pre:= ""
		return pre
	}else {
		minlen := len(strs[0])
		pre = strs[0]
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= minlen {
				minlen = len(strs[i])
				a = i
				fmt.Printf("the a is%v",a)
			}
			//return a
		}
	var j int
	for j = 0; j < minlen; j++ {
		for i := 1; i < len(strs); i++ {
			if strs[i][j] != strs[a][j]  {			
				pre = strs[0][0:j-1] 	
				fmt.Printf("the pre is%v",pre)
			}else{
				pre = strs[0][:j]
				fmt.Printf("the pre2 is%v",pre)
			}
		}
		//return pre
	}
}
	return pre
}
func main() {
	a := []string{"d","d"}
	c := longestCommonPrefix(a)
	fmt.Printf("the res is %v", c)
}
