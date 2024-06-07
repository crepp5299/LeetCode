func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l <= r {
        if l > r {
            break
        }
        m := (l+r)/2
        if nums[m] == target {
            return m
        }
        if nums[m] > target {
            r = m - 1
        } else if nums[m] < target {
            l  = m + 1
        }
    }
    return -1
}