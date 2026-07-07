func isAnagram(s string, t string) bool {
   cmap := make(map[rune]int)
   if len(t) != len(s){
      return false
   }
   for _,i := range s {
      cmap[i] = cmap[i]+1
   }
   for _,i := range t {
      cmap[i] = cmap[i]-1
      if cmap[i] < 0 {
         return false
      }
   }
   return true
}
