func hasDuplicate(nums []int) bool {
    newMap := make(map[int]struct{})

	for _,v := range(nums){
		newMap[v] = struct{}{}
	}

	if len(newMap) != len(nums){
		return true
	}
	return false

}
