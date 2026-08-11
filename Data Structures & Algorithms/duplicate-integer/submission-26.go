func hasDuplicate(nums []int) bool {
	seen := map[int]bool{}

	for _, value := range(nums) {
		seen[value] = true
	}
	return len(seen) < len(nums)
}
