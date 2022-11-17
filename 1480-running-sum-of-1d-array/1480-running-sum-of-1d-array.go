func sum(s []int, ch chan int) {
    total := 0
    for i := 0; i < len(s); i++ {
        total += s[i]
    }
    ch <- total // Send total to channel ch.
} 

func runningSum(nums []int) []int {
    result := [] int{}
    
    ch := make(chan int)
    
    for i := 0; i < len(nums); i++ {
        go sum(nums[:i+1], ch)
        result = append(result, <-ch) // append ch value to result array
    }
    
    close(ch) // close channel ch
    
    return result
    
}