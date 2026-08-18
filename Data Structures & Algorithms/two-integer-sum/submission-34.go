func twoSum(nums []int, target int) []int {
    valuesMap := map[int]int{}

    for i,v := range(nums){
        valuesMap[v]=i
    }
    for index,val := range(nums){
        complement := target-val
        if complementIndex, exists:= valuesMap[complement]; exists{
            if complementIndex>index {
            return []int{index, complementIndex}}
        }
    }

return []int{}
}