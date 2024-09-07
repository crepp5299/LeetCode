func numJewelsInStones(jewels string, stones string) int {
    m := make(map[string]int)
    for i := range stones {
        m[string(stones[i])]++
    }
    result := 0
    for i := range jewels {
        result+=m[string(jewels[i])]
    }
    return result
}