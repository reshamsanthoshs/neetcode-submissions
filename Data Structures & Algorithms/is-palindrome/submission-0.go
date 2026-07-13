func isPalindrome(s string) bool {
	runes := []rune(s)
	left := 0
	right := len(runes)-1

	for left < right {
		
		if !(unicode.IsLetter(runes[left]) || unicode.IsDigit(runes[left]))  {
			left++
			continue
		}
		if !(unicode.IsLetter(runes[right]) ||unicode.IsDigit(runes[right])){
			right--
			continue
		}

		if unicode.ToLower(runes[left]) == unicode.ToLower(runes[right]){
			left++
			right--
			continue
		} else{
			return false
		}
	}

	return true
	
	
}
