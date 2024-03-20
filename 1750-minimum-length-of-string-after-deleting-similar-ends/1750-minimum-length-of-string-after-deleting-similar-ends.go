func minimumLength(s string) int {
    begin := 0
    end := len(s) - 1

    for begin < end && s[begin] == s[end] {
        c := s[begin]
        for begin <= end && s[begin] == c {
            begin++
        }
        for begin < end && s[end] == c {
            end--
        }
    }
    return end-begin+1
}

