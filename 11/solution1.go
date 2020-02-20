package solution1
import("math")
func maxArea(height []int) int {
    var res,h int
    i:=0
    j:=len(height)-1
    if i ==j{
        return 0
    }
    for i<j{
         h = int(math.Min(float64(height[i]),float64(height[j])))
         
        res = int(math.Max(float64(res), float64(h*(j-i))))
        
        if height[i]<height[j]{
            i++
        }else{
            j--
        }
        
    }
    return res
}