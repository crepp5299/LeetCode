func minimumLength(s string) int {
    left := 0
    n:= len(s) - 1
    right := n
    c := byte('d')
    for left < n && right >= 0 && left <= right {
        if left == right {
            if s[left] != c {
                return 1
            }
            return 0
        }
        if s[right] == c {
            right--
            continue
        }
        if s[left] == c {
            left++
            continue
        }
        if s[left] != s[right] && c == 'd' {
            break
        }
        if s[left] == s[right] {
            c = s[left]
            right--
            continue
        }
        if s[left] != s[right] {
            c = 'd'
            continue
        }
    }
    return len(s[left:right+1])
}

