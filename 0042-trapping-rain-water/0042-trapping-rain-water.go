func trap(height []int) int {
    h := len(height)
    l, maxL, maxR := make([]int, h), 0, 0
    for i := range height {
        l[i] = maxL
        maxL = max(height[i], maxL)
    }
    maxL = 0
    for i := range height {
        if min(l[h-1-i], maxR) - height[h-1-i] > 0 {
            maxL += min(l[h-1-i], maxR) - height[h-1-i]
        }
        maxR = max(height[h-1-i], maxR)
    }
    return maxL
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}