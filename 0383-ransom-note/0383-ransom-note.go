func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[byte]int, 26)
    for i := range magazine {
        m[magazine[i]]++
    }
    for i := range ransomNote {
        if m[ransomNote[i]] > 0 {
            m[ransomNote[i]]--
        } else {
            return false
        }
    }
    return true
}