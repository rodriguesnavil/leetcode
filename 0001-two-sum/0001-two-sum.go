func twoSum(nums []int, target int) []int {

    //     Brute force approach, I will be writing optimal approach sooooon
    result := []int{}
    for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
                result = append(result, i, j)
                return result
			}
		}
	}
    return []int{}
}