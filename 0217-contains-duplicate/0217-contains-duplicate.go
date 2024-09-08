func containsDuplicate(nums []int) bool {
    n := len(nums)
    m := make(map[int]bool, n)
    for i := 0; i<n; i++ {
        if m[nums[i]] {
            return true
        }
        m[nums[i]] = true
    }
    return false
}