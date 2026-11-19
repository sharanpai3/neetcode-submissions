func twoSum(nums []int, target int) []int {
	newMap:= make(map[int]int)

	for i,v:=range(nums){
		newMap[v] = i
	}

	for in, va := range(nums){
		diff := target - va
		if _, exists:= newMap[diff]; exists && in < newMap[diff] {
			return []int{in, newMap[diff]}}
	}
	return []int{0,0}
}
