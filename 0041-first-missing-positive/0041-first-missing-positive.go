func firstMissingPositive(nums []int) int {
    sort.Ints(nums)
    num := 1
    prevNum := 0
    for _, val := range nums {
        if val <= 0 {
            continue
        }
        if val == prevNum {
            continue
        }
        if val == num {
            prevNum = val
            num ++
        } else {
            break
        }
    }
    return num
}