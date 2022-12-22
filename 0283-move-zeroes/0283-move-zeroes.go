func moveZeroes(nums []int)  {
    i, j := 0, 0
    
    for i < len(nums) {
        if findAbs(nums[i]) > 0 {
            nums[i], nums[j] = nums[j], nums[i]
            j++
        }
        i++
    }
}


func findAbs(num int) int {
    if num > 0 {
        return num
    } else {
        return - num
    }
}