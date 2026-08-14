func twoSum(nums []int, target int) []int {
    valuesMap := map[int]int{}

    for i,v := range(nums){
        diff := target-v
        if j, exists:= valuesMap[diff]; exists {
            return []int{j,i}
        }
        valuesMap[v] = i 
    }

return []int{}
}