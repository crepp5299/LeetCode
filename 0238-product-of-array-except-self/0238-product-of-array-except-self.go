func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    x := 1
    for i := range nums {
        result[i] = x
        x *= nums[i]
    }
    x = 1
    for n > 0 {
        result[n-1] = x*result[n-1]
        x *= nums[n-1]
        n--
    }
    return result
}