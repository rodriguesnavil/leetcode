func minimumAbsDifference(arr []int) [][]int {
    if len(arr) < 2 {
        return [][] int{}
    }
    
    var result [][]int
    
    min := math.MaxInt
    
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    
    for i := 1; i < len(arr); i++ {
        if arr[i] - arr[i-1] < min {
            min = arr[i] - arr[i-1]
        }
    }
    
    for i := 1; i < len(arr); i++ {
        if arr[i] - arr[i-1] == min {
            result = append(result, []int{arr[i-1], arr[i]})
        }
    }
    
    return result
}