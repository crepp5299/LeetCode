func findDuplicates(nums []int) []int {
    n := len(nums)
    if n == 0 {
        return []int{}
    }
    a := make([]int, n+1)
    for _, val := range nums {
        a[val]++
    }
    result := []int{}
    
    for i, e := range a {
        if e == 2 {
            result = append(result, i)
        }
    }
    return result
}