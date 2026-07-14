func twoSum(numbers []int, target int) []int {
	  seen := make(map[int]int)
      for i := 0 ; i < len(numbers) ; i ++ {
		comp := target - numbers[i]
		if val,ok := seen[comp];ok{
           return []int{val,i+1}
		}
		seen[numbers[i]] = i+1
	  }
	return nil
}     
