package solution1
import("math")
func maxArea(height []int) int {
    var res,h int
    i:=0
    j:=len(height)-1
    for i<j{
         h = int(math.Min(float64(height[i]),float64(height[j])))
         
        res = int(math.Max(float64(res), float64(h*(j-1))))
        
        if height[i]<height[j]{
            i++
        }else{
            j--
        }
        
    }
    return res
}