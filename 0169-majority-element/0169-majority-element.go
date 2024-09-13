func majorityElement(nums []int) int {
    m := make(map[int]int)
    n := len(nums)
    for i := 0; i < n; i++ {
        m[nums[i]]++
    }
    for k, v := range m {
        if v > n/2 {
            return k
        }
    }
    return 0
}