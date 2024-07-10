func minOperations(logs []string) int {
    var result int
    for i := range logs {
        if logs[i] == "./"{
            continue
        } 
        if logs[i] != "../"{
            result++
        } else if logs[i] == "../" && result > 0 {
            result--
        }
    }
    return result 
}