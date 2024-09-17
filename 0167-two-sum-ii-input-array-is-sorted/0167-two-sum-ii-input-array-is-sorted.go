func twoSum(numbers []int, target int) []int {
    i, j := 0, len(numbers) - 1
    for i < j {
        if a := numbers[i] + numbers[j]; a == target {
            return []int{i+1, j+1}
        } else if a > target {
            j--
        } else {
            i++
        }
    }
    return []int{0, 0}
}