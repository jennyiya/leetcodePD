package main
import(
	"fmt"
)

func searchInsert(nums []int, target int) int {
    var i int
    if len(nums)==0 {
        return 0
	}
    if  target <= nums[0]{
        return  0
    }
	if target > nums[len(nums)-1]{
        return len(nums)
	}
    if target == nums[len(nums)-1]{
        return len(nums)-1
	}
	var j int
    for i:=0;i<len(nums);i++{
          if target > nums[i]{
		    j++
	       }else{		   
				return i   
	       }
	    
	}
	return i
}
func main(){
	a := []int{1,3,5,6}
	res := searchInsert(a,2)
	fmt.Printf("the res is %v",res)
}
 
