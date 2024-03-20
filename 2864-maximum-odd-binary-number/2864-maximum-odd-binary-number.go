func maximumOddBinaryNumber(s string) string {
    result1 := ""
    result2 := ""
    for i:= range s {
        if s[i] == '1'{
            result1+="1"
        } else {
            result2+="0"
        }
    }
    return result1[1:] + result2 + "1"
}
