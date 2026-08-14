func twoSum(nums []int, target int) [2]int {
    for i:=0; i<len(nums); i++{
        for j:=0; j<len(nums); j++{
            if (nums[i]+nums[j]==target && j!=i){
                return [2]int{i, j}
            }
        }
    } 
    return [2]int{}
}