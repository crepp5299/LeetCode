func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[string]int)
    for i := range magazine {
        m[string(magazine[i])]++
    }
    for i := range ransomNote {
        if m[string(ransomNote[i])] > 0 {
            m[string(ransomNote[i])]--
        } else {
            return false
        }
    }
    return true
}