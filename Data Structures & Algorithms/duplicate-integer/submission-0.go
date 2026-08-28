func hasDuplicate(nums []int) bool {
    hash_map := map[int]int{}
    for _,num := range nums {
        hash_map[num] += 1
        if hash_map[num] > 1{
            return true
        }
    
}
return false
}