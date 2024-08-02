func minSwaps(nums []int) int {
    n := len(nums)
    nums2 := make([]int, 0, n)
    nums2 = append(nums2, nums...)
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
        if nums2[(i+len1)%n] == 1 && nums2[i] == 0 {
            count--
        } else if nums2[(i+len1)%n] == 0 && nums2[i] == 1 {
            count++
        }
        if min > count {
            min = count
        }
    }
    return min
}