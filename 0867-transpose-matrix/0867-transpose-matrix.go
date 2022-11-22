func transpose(matrix [][]int) [][]int {
    m, n := len(matrix), len(matrix[0])
    
    res := make([][]int, n)
    
    for i := range res {
        res[i] = make([]int, m)
    }
    
    for r := 0; r < n; r++ {
        for c := 0; c < m; c++ {
            res[r][c] = matrix[c][r]
        }
    }
    
    return res
}