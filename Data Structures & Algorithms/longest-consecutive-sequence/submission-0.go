func longestConsecutive(nums []int) int {
	// Step 1: Put every number into a set
	seen := make(map[int]bool)
	for _, num := range nums {
		seen[num] = true
	}

	maxLength := 0
	for _, num := range nums {
		if seen[num-1] {
			continue
		}

		current := num
		length := 1
		for seen[current+1] {
			current++
			length++
		}

		// Step 4: Update answer
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
