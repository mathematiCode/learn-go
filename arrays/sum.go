package arrays

func Sum3(nums [3]int) int {
	var total int
	for i := 0; i < 3; i++ {
		total += nums[i]
	}
	return total
}

func Sum(nums []int) int {
	var total int
	for _, item := range nums {
		total += item
	}
	return total
}
