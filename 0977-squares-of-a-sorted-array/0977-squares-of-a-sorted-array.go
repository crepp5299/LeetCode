func sortedSquares(nums []int) []int {
    n := len(nums)
    i := 0
    j := n - 1
    result := make([]int, n)
    for {
        if n == 0 {
            break
        }
        if nums[i] < 0 && -nums[i] >= nums[j]{
            result[n - 1] = nums[i] * nums[i]
            i++
        } else {
            result[n - 1] = nums[j] * nums[j]
            j--
        }
        n--
    }
    return result
}