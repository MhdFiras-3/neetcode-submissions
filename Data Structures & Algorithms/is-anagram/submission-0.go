func isAnagram(s string, t string) bool {
		if len(s) != len(t) {
		return false
	}
	sMap := map[rune]int{}
	tMap := map[rune]int{}

	for i, char := range s{
		sMap[char]++
		tMap[rune(t[i])]++
	}
	for _, char := range t{
		if sMap[char] != tMap[char] {
			return false
		}
	}
	return true
}
