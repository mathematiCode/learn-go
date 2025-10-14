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

func SumAll(arrays ...[]int) []int {
	var result []int
	for _, array := range arrays {
		var total int
		for _, number := range array {
			total += number
		}
		result = append(result, total)
	}
	

hi := 5
for _, hi > 0 {
	hi = hi - 1
	fmt.Printf("Hi %d", hi)
}

return result
}