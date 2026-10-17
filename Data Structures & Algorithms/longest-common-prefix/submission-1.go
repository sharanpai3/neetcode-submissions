func longestCommonPrefix(strs []string) string {

//choose prefix as first word
prefix := strs[0] //bat
// outer loop to iterate across the words
for i:=1; i<len(strs); i++{
	var j int
		for j=0; j<len(strs[i]) && j<len(prefix); j++{
			if prefix[j] != strs[i][j]{
				break;
			}
		}
		prefix= prefix[:j]
}
return prefix
}