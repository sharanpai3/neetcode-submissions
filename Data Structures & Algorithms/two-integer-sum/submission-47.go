func twoSum(nums []int, target int) []int {
    newMap := make(map[int]int)

	for i,v := range(nums){
		newMap[v] = i //map value to index; value to position
	}

	for indexArr, valueArr := range(nums){
		diff:=target-valueArr
		if j,ok:= newMap[diff]; ok && indexArr<j{
			return []int{indexArr,j}
		}
	}
	
	return []int{0,0}

}
