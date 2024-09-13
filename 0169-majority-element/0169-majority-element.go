// func majorityElement(nums []int) int {
//     m := make(map[int]int)
//     n := len(nums)
    
//     for i := 0; i < n; i++ {
//         m[nums[i]]++
//     }
//     for k, v := range m {
//         if v > n/2 {
//             return k
//         }
//     }
//     return 0
// }

func majorityElement(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
    })
    count := 1
    n := len(nums)
    for i := 1; i < n; i++ {
        if nums[i] == nums[i-1] {
            count++
        } else {
            count = 1
        }
        if count > n/2 {
            return nums[i]
        }
    }
    return nums[0]
}