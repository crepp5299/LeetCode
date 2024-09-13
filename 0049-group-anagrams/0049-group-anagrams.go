func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    for i := 0; i < len(strs); i++ {
        s := []rune(strs[i])
        sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
        m[string(s)] = append(m[string(s)], strs[i])
    }
    rs := make([][]string, 0)
    for _, val := range m {
        rs = append(rs, val)
    }
    return rs
}

