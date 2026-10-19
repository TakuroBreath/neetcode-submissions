func search(nums []int, target int) int {
	return sh(nums, target, 0, len(nums)-1)
}

func sh(nums []int, t, l, r int) int {
	if l > r {
		return -1
	}
	
	m := (l+r)/2

	if nums[m] == t {
		return m
	} else if nums[m] > t {
		return sh(nums, t, l, m-1)
	} else {
		return sh(nums, t, m+1, r)
	}
}