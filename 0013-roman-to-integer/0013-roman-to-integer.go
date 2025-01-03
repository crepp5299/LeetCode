func romanToInt(s string) int {
    sum := 0
	rim := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
    sum += rim[string(s[0])]
    

    for i := 1; i < len(s); i++ {
		sum += rim[string(s[i])]
        if rim[string(s[i-1])] < rim[string(s[i])] {
            sum -= 2 * rim[string(s[i-1])]
        }
	}

	return sum
}