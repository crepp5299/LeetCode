func maxArea(height []int) int {
    i, j := 0, len(height) - 1
    max := 0
    for i < j {
        if min(height[i], height[j]) * (j-i) > max {
            max = min(height[i], height[j]) * (j-i)
        }
        if height[i] <= height[j] {
            i++
            continue
        } 
        if height[j] <= height[i]{
            j--
        }
    }
    return max
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}