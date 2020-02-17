package solution2

func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)
	var res, i int
	for i = 0; i < len(nums); i++ {
		temp[i] = nums[i]
    }
	for j := 0; j < len(nums); j++ {
			res = target - nums[j]
			for k, v := range temp {
				if v == res && j != k{
					return []int{j, k}
				}			
		}
	}
	return nil
}
