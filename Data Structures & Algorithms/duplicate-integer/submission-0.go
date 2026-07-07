func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)
	l := len(nums)
	for i := 0; i < l ; i++ {

		if seen[nums[i]] {
			return true
		} else {
			seen[nums[i]] = true
		}
		
	}
	return false
}
