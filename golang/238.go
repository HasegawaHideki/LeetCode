func productExceptSelf(nums []int) []int {
	size := len(nums)
	var answer []int = make([]int, size)

	mul := 1
	for i := 0; i < size; i++ {
		answer[i] = mul
		mul *= nums[i]
	}
	mul = 1
	for i := size - 1; i >= 0; i-- {
		answer[i] *= mul
		mul *= nums[i]
	}

	return answer
}