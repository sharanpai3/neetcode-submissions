func removeElement(nums []int, val int) int {
	var k int
    for _,v:= range(nums){
		if v != val{
			nums[k] = v
			k++
		}
	}
	return k
}
