func twoSum(nums []int, target int) []int {

	for i := len(nums) - 1; i == 0; i-- {
		if nums[i] > target {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}

	result := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
			}
		}
	}

	return result
}
