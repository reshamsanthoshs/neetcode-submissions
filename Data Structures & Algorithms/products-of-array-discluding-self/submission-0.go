func productExceptSelf(nums []int) []int {
	l := len(nums)
	lp := make([]int, l)
	rp := make([]int, l)
	m := make([]int,l)
	left := 1
	right := 1

    for i:=0; i < len(nums) ; i ++ {
		lp[i] = left
		left *= nums[i]
	}
	for j:= len(nums)-1 ; j >=0; j--{
		rp[j] = right
		right *= nums[j]
	}
	for k:=0 ; k < len(nums) ; k ++{
		m[k] = rp[k] * lp[k]
	}
	return m
}
