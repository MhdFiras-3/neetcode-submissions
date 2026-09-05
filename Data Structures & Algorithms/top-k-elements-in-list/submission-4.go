	func topKFrequent(nums []int, k int) []int {
		mapCount := make(map[int]int)

		for _, num := range nums {
			mapCount[num]++
		}
		freqSlice := make([][]int, len(nums) + 1)
		for num,freq := range mapCount {
			freqSlice[freq] = append(freqSlice[freq],num)
		}

		result := []int{}
		for i := len(freqSlice) - 1; i > 0; i-- {
			for _, num := range freqSlice[i] {
				result = append(result, num)
				if len(result) == k {
					return result
				}

			}
		}


	return result


}
