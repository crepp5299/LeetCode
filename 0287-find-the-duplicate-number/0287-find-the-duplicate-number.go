func findDuplicate(nums []int) int {
    for _, val := range nums {
        num := val
        if num < 0 {
            num = 0 - val
        }
        if nums[num] < 0 {
            return num
        }
        nums[num]*=-1
    }
    return 0
}