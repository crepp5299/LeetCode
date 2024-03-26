func firstMissingPositive(nums []int) int {
    n := len(nums)
    for i := range nums {
        if nums[i] <=0 || nums[i] > n {
            nums[i] = n + 1
        }
    }
    for _, val := range nums {
        x := abs(val)
        if x == n + 1 {
            continue
        }
        if nums[x-1] > 0 {
            nums[x-1]*=-1
        }
    }
    for i := range nums {
        if nums[i] > 0  {
            return i + 1
        }
    }
    return n + 1
}

func abs(num int) int {
    if num < 0 {
        return 0 - num
    }
    return num
}