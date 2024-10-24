func trap(height []int) int {
    h := len(height)
    l, r, maxL, maxR, rs := make([]int, h), make([]int, h), 0, 0, 0
    for i := range height {
        l[i] = maxL
        r[h-1-i] = maxR
        maxL = max(height[i], maxL)
        maxR = max(height[h-1-i], maxR)
    }
    for i := range height {
        if min(l[i], r[i]) - height[i] > 0 {
            rs += min(l[i], r[i]) - height[i]
        }
    }
    return rs
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