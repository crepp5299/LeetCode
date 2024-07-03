func minDifference(nums []int) int {
    l:= len(nums)
    if l <= 4 {
        return 0
    }
    sort.Ints(nums)
    t:=l-3-1
    rs := math.MaxInt
    for i:= 0;i+t < l;i++ {
        if nums[i+t] - nums[i] < rs {
            rs = nums[i+t] - nums[i]
        }
    }
    return rs
}