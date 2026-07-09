func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	for i, num := range nums {
		compliment := target - nums[i]
		if j,ok := seen[compliment];ok {
            return []int{j,i}
		}
		seen[num] = i 
	}
	return []int{}

}
