func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i:=1; i<len(strs);i++{
		var j int
		for j=0;j<len(prefix) && j<len(strs[i]); j++{
			if prefix[j] != strs[i][j]{
				break
			}
		}
			prefix = prefix[:j]
	}
	return prefix
}