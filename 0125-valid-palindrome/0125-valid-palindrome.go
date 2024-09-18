func isPalindrome(s string) bool {
    s = strings.ToUpper(s)
    i,j:=0,len(s)-1
    for i < j {
        if !checkOk(s[i]) {
            i++
        } else if !checkOk(s[j]) {
            j--
        } else if s[i]==s[j] {
            i++
            j--
        } else {
            return false
        }
    }
    return true
}

func checkOk(s byte) bool {
    if s >= 65 && s <= 90 {
        return true
    }
    if s >= 48 && s <= 57 {
        return true
    }
    return false
}