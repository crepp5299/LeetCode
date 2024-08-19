func isSubsequence(s string, t string) bool {
    if s == "" {
        return true
    }
    s1 := []rune(s)
    t1 := []rune(t)
    l1 := len(s1)
    count := l1
    for i := 0; i < len(t1); i++ {
        if s1[l1-count] == t1[i] {
            count--
        }
        if count == 0 {
            return true
        }
    }
    return false
}