func isValidSudoku(board [][]byte) bool {
    for i := 0; i < 9; i++ {
        m := make(map[byte]bool)
        for j := 0; j < 9; j++ {
            if m[board[i][j]] {
                return false
            }
            if board[i][j] != '.' {
                m[board[i][j]] = true
            }
        }
    }
    
    for i := 0; i < 9; i++ {
        m := make(map[byte]bool)
        for j := 0; j < 9; j++ {
            if m[board[j][i]] {
                return false
            }
            if board[j][i] != '.' {
                m[board[j][i]] = true
            }
        }
    }
    items := [][]int{
        {0, 0}, {0, 3}, {0, 6},
        {3, 0}, {3, 3}, {3, 6},
        {6, 0}, {6, 3}, {6, 6},
    }
    for index := range items {
        m := make(map[byte]bool)
        for i := 0; i < 3; i++ {
            for j := 0; j < 3; j++ {
                if m[board[items[index][0]+i][items[index][1]+j]] {
                return false
            }
            if board[items[index][0]+i][items[index][1]+j] != '.' {
                m[board[items[index][0]+i][items[index][1]+j]] = true
            }
        }
        }
    }
    
    return true
}