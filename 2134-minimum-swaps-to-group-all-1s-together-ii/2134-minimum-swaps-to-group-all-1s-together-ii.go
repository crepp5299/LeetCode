func minSwaps(nums []int) int {
    n := len(nums)
    len1 := 0
    for i := range nums {
        if nums[i] == 1 {
            len1++
        }
    }
    count := 0
    for i:= 0; i < len1; i++ {
        if nums[i] == 0 {
            count++
        }
    }
    min := count
    for i := 0; i < n; i++ {
        if nums[(i+len1)%n] == 1 && nums[i] == 0 {
            count--
        } else if nums[(i+len1)%n] == 0 && nums[i] == 1 {
            count++
        }
        if min > count {
            min = count
        }
    }
    return min
}