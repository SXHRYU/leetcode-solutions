package leetcode

func pivotIndex(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}
	var sum, leftSum int
	for _, num := range nums {
		sum += num
	}
	for index, num := range nums {
		if leftSum*2+num == sum {
			return index
		}
		leftSum += num
	}
	return -1
}
