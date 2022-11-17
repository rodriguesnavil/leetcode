func runningSum(nums []int) []int {
//     optimized solution
    for i := 1; i < len(nums); i++ {
        nums[i] = nums[i] + nums[i -1]
    }
    return nums
}