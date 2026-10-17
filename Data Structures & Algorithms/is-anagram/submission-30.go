func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	sMap := make(map[rune]int)
	for _,v := range(s){
		sMap[v] = sMap[v] + 1
	}

	tMap := make(map[rune]int)
	for _,val := range(t){
		tMap[val] = tMap[val] + 1
	}

	for index, value := range(sMap){
		if tMap[index] != value{
			return false
		}
	}
	return true

}
