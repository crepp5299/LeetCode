func reverseString(s []byte)  {
    i, n := 0, len(s)
    for i < n/2 {
        s[i], s[n-i-1] = s[n-i-1], s[i]
        i++
    }
}