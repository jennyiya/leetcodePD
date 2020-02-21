package main
import("fmt")
func removeElement(nums []int, val int) int {
	j := 0
	if len(nums)==0{
		return 0
	}
	var res []int
    for i:=0;i<len(nums);i++{
        if nums[i] != val{
		   //res[j] = nums[i]
		   res = append(res,nums[i])
           j++
        }
	}
	fmt.Printf("res is%v",res)
    return j
}

func main(){
	//var res []int
	a :=[]int{0,1,2,2,3,0,4,2}
	l := removeElement(a,2)
	fmt.Printf("the arr is %v,the len is %v",a,l)
}