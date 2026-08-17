type NumArray struct {
    prefix []int
}

func Constructor(nums []int) NumArray {
    prefix := make([]int, len(nums))
    cur := 0
    for i, num := range nums {
        cur += num
        prefix[i] = cur
    }
    return NumArray{prefix: prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
    rightSum := this.prefix[right]
    leftSum := 0
    if left > 0 {
        leftSum = this.prefix[left-1]
    }
    return rightSum - leftSum
}