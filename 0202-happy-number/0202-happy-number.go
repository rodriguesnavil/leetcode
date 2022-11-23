func isHappy(n int) bool {
    isRepeated := make(map[int]bool)
    
    for n != 1 {
        current := n
        sum := 0
        for current != 0 {
            sum += (current % 10) * (current % 10)
            current /= 10
        }
        
        if isRepeated[sum] {
            return false
        }
        
        isRepeated[sum] = true
        
        n = sum
    }
    return true
}