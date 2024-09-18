func isPalindrome(s string) bool {
    s = strings.ToUpper(s)
    i,j:=0,len(s)-1
    for i < j {
        if !(s[i] >= 65 && s[i] <= 90 || s[i] >= 48 && s[i] <= 57) {
            i++
        } else if !(s[j] >= 65 && s[j] <= 90 || s[j] >= 48 && s[j] <= 57) {
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
