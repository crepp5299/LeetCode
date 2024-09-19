func threeSum(nums []int) [][]int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] <nums[j]
    })
    rs := make([][]int, 0)
    n := len(nums)
    for index := 0; index < n; index++ {
        if index > 0 && nums[index] == nums[index-1]{
            continue
        }
        i, j := index + 1, n - 1
        for i < j {
            if i > index + 1 && nums[i] == nums[i-1] {
                i++
            } else if j < n - 1 && nums[j] == nums[j + 1] {
                j--
            } else if nums[i] + nums[j] + nums[index] == 0 {
                rs = append(rs, []int{nums[index], nums[i],  nums[j]})
                i++
                j--
            } else if nums[i] + nums[j] + nums[index] < 0 {
                i++
            } else {
                j--
            }
        }
    }
    return rs
}