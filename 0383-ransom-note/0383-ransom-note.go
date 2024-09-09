func canConstruct(ransomNote string, magazine string) bool {
	h := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		h[magazine[i]-97]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if h[ransomNote[i]-97] <= 0 {
			return false
		}
		h[ransomNote[i]-97]--
	}
	return true
}