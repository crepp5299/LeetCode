func minDifference(nums []int) int {
    sort.Ints(nums)
    l:= len(nums)
    if l <= 4 {
        return 0
    }
    t:=l-3-1
    rs := nums[t] - nums[0]
    for i:=1;i+t < l;i++ {
        if nums[i+t] - nums[i] < rs {
            rs = nums[i+t] - nums[i]
        }
    }
    return rs
}