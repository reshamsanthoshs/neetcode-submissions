func groupAnagrams(strs []string) [][]string {
   groups := make(map[[26]int][]string)
   for _, word := range strs {
        var count [26]int
        for _, ch := range word {
            count[ch-'a']++
			
        }
        groups[count] = append(groups[count], word)
		print(groups[count])
    }
    result := [][]string{}

    for _, group := range groups {
        result = append(result, group)
    }
    return result
}


