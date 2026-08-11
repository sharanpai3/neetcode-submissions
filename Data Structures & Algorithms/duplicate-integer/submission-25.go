func hasDuplicate(nums []int) bool {
	seen := map[int]bool{}

	for _, value := range(nums) {
		seen[value] = true
	}
	if len(seen) < len(nums){
	return true
	}else{
		return false
	}
}
