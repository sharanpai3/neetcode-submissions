func hasDuplicate(nums []int) bool {
	seen := map[int]bool{}

	for _, value := range(nums) {
		if seen[value] == true{
			return true
		}
		seen[value] = true
	}
	return false
}
