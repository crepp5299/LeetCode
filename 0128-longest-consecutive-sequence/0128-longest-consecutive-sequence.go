func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    max, count:= 0, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] + 1 {
            count++
        } else if nums[i] > nums[i-1] + 1{
            if max < count {
                max = count
            }
            count = 1
        }
    }
    if max < count {
        max = count
    }
    return max
}