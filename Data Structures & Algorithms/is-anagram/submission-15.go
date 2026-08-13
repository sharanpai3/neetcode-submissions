
func isAnagram(s string, t string) bool {
	if (len(s) != len(t)){
		return false
	}

	sRunes := []rune(s)
	tRunes := []rune(t)

	sFreq := map[rune]int{}
	tFreq := map[rune]int{}

	for _, valuei :=range(sRunes){
		sFreq[valuei]++
	}
		for _, valuek :=range(tRunes){
		tFreq[valuek]++
	}

	for _,value:=range(sRunes){
		if sFreq[value] != tFreq[value]{
			return false
		}
	}
return true
}
