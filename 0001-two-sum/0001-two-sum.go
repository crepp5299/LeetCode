func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i, e := range nums {
        if m[e]  > 0 {
            return []int{m[e] - 1, i}
        }
        m[target - e] = i + 1
    }
    return []int{0 , 0}
}