func maxNumberOfBalloons(text string) int {
    b := map[rune]int{'b': 1, 'a' : 1, 'l' : 2, 'o' : 2, 'n':1 }
    m := make(map[rune]int, 26)
    for _, c := range text {
        m[c]++
    }
    min := math.MaxInt
    for k, v := range b {
        if m[k]/v < min {
            min = m[k]/v
        }
    }
    return min
}