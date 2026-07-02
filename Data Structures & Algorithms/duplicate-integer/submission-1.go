func hasDuplicate(nums []int) bool {
    // Given unsorted array

	presenceMap:= make(map[int]struct{})

	for _, num:= range nums {

		if _, ok:= presenceMap[num]; ok{
			return true
		}
		presenceMap[num] = struct{}{}

	}  
	return false
}
