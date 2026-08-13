
func isAnagram(s string, t string) bool {
	if (len(s) != len(t)){
		return false
	}

	sRunes, tRunes := map[rune]int{},  map[rune]int{}
	for i, ch := range(s){ //frequency map
		sRunes[ch]++
		tRunes[rune(t[i])]++
	}
	for k,v := range(sRunes){
		if (tRunes[k] != v){
			return false
		}
	}
	return true
}
