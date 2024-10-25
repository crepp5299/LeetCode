func trap(height []int) int {
    i, j := 0, len(height) - 1
    maxL, maxR, rs := 0, 0, 0
    for i < j {
        if height[i] <= height[j]{
            if maxL - height[i] > 0 {
                rs += maxL - height[i]
            }
            if maxL < height[i] {
                maxL = height[i]
            }
            i++
        } else {
            if maxR - height[j] > 0 {
                rs += maxR - height[j]
            }
            if maxR < height[j] {
                maxR = height[j]
            }
            j--
        }
    }
    return rs
}
