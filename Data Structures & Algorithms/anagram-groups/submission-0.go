func groupAnagrams(strs []string) [][]string {
	wordMap := map[string][]string{}
	result := [][]string{}

	for _, word := range strs {
		chars := []byte(word)
		sort.Slice(chars, func(i,j int)bool{
			return chars[i] < chars[j]
		})
		wordMap[string(chars)] = append(wordMap[string(chars)], word)
	}

	for _, value := range wordMap {
		result = append(result, value)
	}
	return result
}
