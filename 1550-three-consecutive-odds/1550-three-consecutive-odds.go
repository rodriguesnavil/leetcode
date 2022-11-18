func threeConsecutiveOdds(arr []int) bool {
    counter := 0
    
    for i := 0; i < len(arr); i++ {
        if arr[i] % 2 != 0 {
            counter += 1
        } else {
            counter = 0
        }
        
        if (counter == 3) {
            return true
        }
    }
    
    return false
}