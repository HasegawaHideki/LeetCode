func moveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			tmp := nums[index]
			nums[index] = nums[i]
			nums[i] = tmp
			index++
		}
	}
}