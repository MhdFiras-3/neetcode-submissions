func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{}

	for i, num := range nums {
		comp := target - num
		if val, ok := numsMap[comp]; ok {
			return []int{val,i}
		}
		numsMap[num] = i
	}
	return []int{}
}
