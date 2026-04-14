package leetcode

import (
	"math/rand"
	"sort"
)

func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func findKthLargest(nums []int, k int) int {
	m := len(nums) - k + 1
	return selectSmallest(nums, 0, len(nums)-1, m)
}

func selectSmallest(nums []int, l, r, i int) int {
	if l >= r {
		return nums[l]
	}
	q := partition(nums, l, r)
	k := q - l + 1
	if k == i {
		return nums[q]
	}
	if i < k {
		return selectSmallest(nums, l, q-1, i)
	} else {
		return selectSmallest(nums, q+1, r, i-k)
	}
}

func partition(nums []int, l, r int) int {
	k := l + rand.Intn(r-l+1)
	nums[k], nums[r] = nums[r], nums[k]
	i := l - 1

	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

func getLeastNumbers(arr []int, k int) []int {
	return selectSmallest1(arr, 0, len(arr)-1, k)[:k]
}

func selectSmallest1(nums []int, l, r, i int) []int {
	if l >= r {
		return nums
	}
	q := partition(nums, l, r)
	k := q - l + 1
	if k == i {
		return nums
	}
	if i < k {
		return selectSmallest1(nums, l, q-1, i)
	} else {
		return selectSmallest1(nums, q+1, r, i-k)
	}
}
