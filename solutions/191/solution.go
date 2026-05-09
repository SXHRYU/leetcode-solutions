package leetcode

import "math/bits"

func hammingWeight(num uint32) int {
	return bits.OnesCount(uint(num))
}

func hammingWeight1(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
