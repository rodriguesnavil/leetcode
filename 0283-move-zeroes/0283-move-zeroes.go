func moveZeroes(nums []int)  {
    j := 0
    
    for i := 0; i < len(nums); i++ {
        if((Abs(nums[i])) > 0) {
            nums[i], nums[j] = nums[j], nums[i]
            j++
        }
    }
}


func Abs(num int) int {
    if (num < 0) {
        return -num
    }
    return num
}