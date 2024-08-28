func merge(intervals [][]int) [][]int {
    sort.Slice(intervals[:], func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    n:= len(intervals)
    result := make([][]int, 0)
    tempItem := []int{intervals[0][0], intervals[0][1]}
    for i := 1; i < n; i++ {
        if tempItem[1] < intervals[i][0] {
            result = append(result, tempItem)
            tempItem = intervals[i]
        } else if tempItem[1] >= intervals[i][0] && tempItem[1] <= intervals[i][1]{
            tempItem[1] = intervals[i][1]
        }
    }
    result = append(result, tempItem)
    return result
}