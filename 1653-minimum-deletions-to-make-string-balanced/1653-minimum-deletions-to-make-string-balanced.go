func minimumDeletions(s string) int {
    runes := []rune(s)
    n := len(runes)
    countA := make([]int, n)
    count := 0
    for i := n - 1; i >= 0; i-- {
        countA[i] = count
        if runes[i] == 'a' {
            count++
        }
    }
    
    count = 0
    
    for i, _ := range runes {
        if countA[i] + count < n {
            n = countA[i] + count
        }
        if runes[i] == 'b' {
            count++
        }
    }
    return n
}