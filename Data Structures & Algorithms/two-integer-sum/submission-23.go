func twoSum(nums []int, target int) []int {
    valuesMap := map[int]int{}

    //form map of value-complement
    for ind,val := range(nums){
        valuesMap[val] = ind
    }

    for i,v := range(nums){
        diff := target-v
        if j, exists:= valuesMap[diff]; exists && j>i {
            return []int{i,j}
        }
    }

return []int{}
}