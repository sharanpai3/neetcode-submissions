
func isAnagram(s string, t string) bool {
	if (len(s) != len(t)){
		return false
	}

	// sRunes, tRunes := []rune(s), []rune(t)

	sMap, tMap := map[rune]int{}, map[rune]int{}

	for i,v := range(s){
		sMap[v]++
		tMap[rune(t[i])]++
	}

	for k, _ := range(sMap){
		if sMap[k] != tMap[k]{
			return false
		}
	}

	return true


}
