
func isAnagram(s string, t string) bool {
	if (len(s) != len(t)){
		return false
	}

	// sRunes, tRunes := []rune(s), []rune(t)

	var count [26]int
	// sMap, tMap := map[rune]int{}, map[rune]int{}

	for i,v := range(s){
		count[v-'a']++
		count[t[i]-'a']--
	}

	for _, value := range(count){
		if value != 0{
			return false
		}
	}

	return true


}
