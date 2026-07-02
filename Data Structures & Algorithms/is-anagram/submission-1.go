func isAnagram(s string, t string) bool {
	// 1. Create a map and insert the count for s

	anagramMap := make(map[rune]int)
	for _, ch := range s {
		anagramMap[ch]++
	}
	// 2. iterate map and decrease the count
	for _, ch := range t {
		anagramMap[ch]--
	}
	// 3. iterate map again and check if the count !=0 return false

	for _, val := range anagramMap {
		if val != 0 {
			return false
		}
	}
	return true
}