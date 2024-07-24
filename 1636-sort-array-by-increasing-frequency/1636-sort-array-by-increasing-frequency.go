func frequencySort(nums []int) []int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })
    arr2 := make([][]int, 101)
    count, val := 1, nums[0]
    for i := 1; i < len(nums); i++{
        if nums[i] != val {
            arr2[count] = append(arr2[count], val)
            count = 1
            val = nums[i]
        } else {
            count++
        }
    }
    arr2[count] = append(arr2[count], val)
    rs := make([]int, 0, 100)
    for i := range arr2 {
        for j := range arr2[i] {
            for k := i; k > 0; k-- {
                rs = append(rs, arr2[i][j])
            }
        }
    }
    return rs
}