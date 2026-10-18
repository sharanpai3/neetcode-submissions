func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	var i int 
	for i=0; i<len(prefix);i++{
		for _,v := range(strs){
			if i==len(v) || v[i]!=prefix[i]{
				return prefix[:i]
			}
		}
	} 
	return prefix
}