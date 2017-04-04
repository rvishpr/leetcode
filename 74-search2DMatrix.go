func searchMatrix(matrix [][]int, target int) bool {
    rows := len(matrix)
    if rows == 0 {
        return false
    }
    cols := len(matrix[0])
    if cols == 0 {
        return false
    }
    
    found := false
    r := 0
    c := cols - 1
    for !found {
        if matrix[r][c] < target {
            r = r + 1
        } else if matrix[r][c] > target {
            c = c - 1
        } else {
            found = true
            break
        }
        if r >= rows || c < 0 {
            break
        }
    }
    
    return found
}
