func averageWaitingTime(customers [][]int) float64 {
    totalWait := customers[0][1]
    lastFinishTime := customers[0][1] + customers[0][0]
    l := len(customers)
    for i := 1; i < l; i++ {
        fmt.Println(lastFinishTime,totalWait)
        if lastFinishTime >= customers[i][0] {
            totalWait += lastFinishTime - customers[i][0] + customers[i][1]
            lastFinishTime += customers[i][1]
        } else {
            totalWait += customers[i][1]
            lastFinishTime = customers[i][1] + customers[i][0]
        }
    }
    fmt.Println(lastFinishTime,totalWait)
    return float64(totalWait)/float64(l)
}