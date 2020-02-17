package solution1
func removeDuplicates(nums []int) int {
    j := 0
    i := 1
    for i < len(nums){
        if nums[i]==nums[j]{
            i++
        }else{
            j++
            nums[j]=nums[i]
            i++
        }
    }
      return j+1     
}