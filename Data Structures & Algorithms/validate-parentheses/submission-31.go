func isValid(s string) bool {
	closeToOpen := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	} 
	stack := []rune{}

	for _,v := range(s){
		if val, exists := closeToOpen[v]; exists{
			if len(stack) > 0 {
				top := stack[len(stack)-1] // assign last element from the stack
				stack = stack[:len(stack)-1]//remove the last element from teh stack
				if top != val {
					return false
				}
			} else {
				return false
			}
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}