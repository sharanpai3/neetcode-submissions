func hasDuplicate(nums []int) bool {
    hasSeen := make(map[int]bool)
	for _,v := range nums{
		if hasSeen[v]{
			return true
		}
		hasSeen[v] = true
	}
	return false
}
