func twoSum(nums []int, target int) []int {
    
	// Create a map to store the presence of the number
	numMap:= make(map[int]int)
	for idx, num:= range nums{
		
			if _, ok:=numMap[target-num]; ok{
				return []int{numMap[target-num], idx}
			}
			numMap[num]= idx
	}
	return []int{}
}
