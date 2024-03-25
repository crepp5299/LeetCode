func findDuplicates(nums []int) []int {
    n := len(nums)
    result := make([]int, 0, n+1)
    for _, val := range nums {
        num := val
        if val < 0 {
            num = 0-val
        }
        if nums[num-1] < 0 {
            result = append(result, num)
        }
        nums[num-1]*=-1
    }
    return result
}