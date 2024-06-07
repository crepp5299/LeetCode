func search(nums []int, target int) int {
    if nums[0] == target {
        return 0
    }
    l, r := 0, len(nums) - 1
    m := (l+r)/2
    for {
        if l > r {
            break
        }
        if nums[m] > target {
            r = m - 1
            m = (l+r)/2
        } else if nums[m] < target {
            l  = m + 1
            m = (l+r)/2
        } else if nums[m] == target {
            return m
        }
    }
    return -1
}