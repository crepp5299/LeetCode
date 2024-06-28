func maximumImportance(n int, roads [][]int) int64 {
    m := make([]int, n)
    for i := range roads {
        m[roads[i][0]]++
        m[roads[i][1]]++
    }
    
    sort.SliceStable(m, func(i, j int) bool {
	    return m[i] > m[j]
    })
    sum := 0
    for i := range m {
        sum += n*m[i]
        n--
    }
    return int64(sum)
}