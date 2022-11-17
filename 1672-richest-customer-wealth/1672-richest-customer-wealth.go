func maximumWealth(accounts [][]int) int {
    maxWealth := 0

    for _, account := range accounts {
        sum := 0
        for _, value := range account {
            sum += value
        }
        if(maxWealth < sum) {
            maxWealth = sum
        }
    }

    return maxWealth
}