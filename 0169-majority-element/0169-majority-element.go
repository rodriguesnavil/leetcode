func majorityElement(nums []int) int {
    majority := make(map[int]int)
    
    for i := 0; i < len(nums); i++ {
        majority[nums[i]] += 1
    }
    
    maxElementValue := math.MinInt
    maxElementKey := 0
    
    for k, v := range majority {
        if v > maxElementValue {
            maxElementValue = v
            maxElementKey = k
        }
    }
    
    return maxElementKey
    
}