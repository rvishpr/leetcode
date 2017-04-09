func minPathSum(grid [][]int) int {
    sum := make([][]int, len(grid))
    rows := len(grid)
    if rows == 0 {
        return 0
    }
    
    cols := len(grid[0])
    if cols == 0 {
        return 0
    }
    
    for i, _ := range sum {
        sum[i] = make([]int, len(grid[0]))
    }
    
    sum[0][0] = grid[0][0]
    for i := 1; i < rows; i++ {
        sum[i][0] = grid[i][0] + sum[i-1][0]
    }
    
    for i := 1; i < cols; i++ {
        sum[0][i] = grid[0][i] + sum[0][i-1]
    }
    
    for i := 1; i < rows; i++ {
        for j := 1; j < cols; j++ {
            ij1 := grid[i][j] + sum[i][j-1]
            i1j := grid[i][j] + sum[i-1][j]
            if ij1 < i1j {
                sum[i][j] = ij1 
            } else {
                sum[i][j] = i1j
            }
        }
    }   
    return sum[rows-1][cols-1]
}
