func maxArea(heights []int) int {
   maxarea := 0
   l := 0
   r := len(heights)-1
   for l<r{
	  area := (r-l)*min(heights[r],heights[l])
	  if area > maxarea{
         maxarea = area
		  
	  }
	  if heights[r] < heights[l]{
		r --
	  }else{
		l ++
	  }

	
   }
   return maxarea
}
