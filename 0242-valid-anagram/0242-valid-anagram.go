// func isAnagram(s string, t string) bool {
//     if len(s) != len(t) {
//         return false
//     }
//     m := make(map[byte]int, 26)
//     for i := 0; i < len(s); i++ {
//         m[s[i]]++
//     }
//     for i := 0; i < len(t); i++ {
//         if m[t[i]] == 0 {
//             return false
//         }
//         m[t[i]]--
//     }
//     return true
// }

func isAnagram(s string, t string) bool {
    s1 := []rune(s)
    t1 := []rune(t)
    sort.Slice(s1, func(i int, j int) bool { return s1[i] < s1[j] })
    sort.Slice(t1, func(i int, j int) bool { return t1[i] < t1[j] })
    return string(s1) == string(t1)
}

    