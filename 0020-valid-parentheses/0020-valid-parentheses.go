func isValid(s string) bool {
    arr := make([]rune, 0)
    for _, c := range s {
        if c == '('  || c == '{' || c == '[' {
            arr = append(arr, c)
        } else if len(arr) > 0 {
            switch c {
                case ')':
                    if arr[len(arr)-1] != '('{
                    return false
                    }
                case '}':
                if arr[len(arr)-1] != '{'{
                    return false
                    }
                case ']':
                    if arr[len(arr)-1] != '['{
                    return false
                    }
            }
            arr = arr[:len(arr)-1]
        } else {
            return false
        }
    }
    if len(arr) > 0 {
        return false
    }
    return true
}