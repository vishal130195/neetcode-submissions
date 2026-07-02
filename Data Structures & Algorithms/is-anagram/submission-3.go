func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	// 1. Create a map and insert the count for s

	anagramMap := make(map[byte]int)
// Single loop using the index!
    for i := 0; i < len(s); i++ {
        anagramMap[s[i]]++ // s[i] pulls the byte at index i
        anagramMap[t[i]]-- // t[i] pulls the byte at index i
    }
	// 3. iterate map again and check if the count !=0 return false

	for _, val := range anagramMap {
		if val != 0 {
			return false
		}
	}
	return true
}