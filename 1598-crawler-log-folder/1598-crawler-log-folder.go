func minOperations(logs []string) int {
    var result int
    a := "./"
    b := "../"
    for i := range logs {
        if logs[i] == a {
            continue
        } 
        if logs[i] != b{
            result++
        } else if logs[i] == b && result > 0 {
            result--
        }
    }
    return result 
}