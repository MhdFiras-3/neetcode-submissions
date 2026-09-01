func topKFrequent(nums []int, k int) []int {
	numsCountMap := make(map[int]int)
	for _, n := range nums {
		numsCountMap[n]++
	}

	numsCountArray := make([][2]int,0,len(numsCountMap))
	for num,count := range numsCountMap {
		numsCountArray = append(numsCountArray, [2]int{num,count})
	}
	sort.Slice(numsCountArray, func(i,j int)bool{
		return numsCountArray[i][1] > numsCountArray[j][1]
	})

	result := make([]int,k)
	for i := 0; i < k; i++ {
		result[i] = numsCountArray[i][0]
	}
	return result

}
