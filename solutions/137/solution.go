package leetcode

func singleNumberII(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}

func singleNumberIIIII(nums []int) int {
	na, nb, nc := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		nb = nb ^ (nums[i] & na)
		na = (na ^ nums[i]) & ^nc
		nc = nc ^ (nums[i] & ^na & ^nb)
	}
	return na & ^nb & ^nc
}

func singleNumberIIIII1(nums []int) int {
	twos, threes, ones := 0xffffffff, 0xffffffff, 0
	for i := 0; i < len(nums); i++ {
		threes = threes ^ (nums[i] & twos)
		twos = (twos ^ nums[i]) & ^ones
		ones = ones ^ (nums[i] & ^twos & ^threes)
	}
	return ones
}
