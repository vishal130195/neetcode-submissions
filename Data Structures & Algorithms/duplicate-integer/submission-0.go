func hasDuplicate(nums []int) bool {
    // Given unsorted array

	presenceMap:= make(map[int]bool)

	for _, num:= range nums {

		if _, ok:= presenceMap[num]; !ok{
			presenceMap[num] = true
		}else{
			return true
		}

	}  
	return false
}
