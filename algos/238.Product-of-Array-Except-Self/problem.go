package algos

func productExceptSelf(nums []int) []int {
	var res, left, right = make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i-1] * left[i-1]
	}
	right[len(right)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}
	for i := 0; i < len(nums); i++ {
		res[i] = left[i] * right[i]
	}
	return res
}

func optimized(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	res[0] = 1
	for i := 1; i < l; i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	rightProduct := 1
	for i := l - 1; i >= 0; i-- {
		res[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return res
}
