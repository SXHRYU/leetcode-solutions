package leetcode

func sumSubarrayMins(A []int) int {
	stack, dp, res, mod := []int{}, make([]int, len(A)+1), 0, 1000000007
	stack = append(stack, -1)

	for i := 0; i < len(A); i++ {
		for stack[len(stack)-1] != -1 && A[i] <= A[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		dp[i+1] = (dp[stack[len(stack)-1]+1] + (i-stack[len(stack)-1])*A[i]) % mod
		stack = append(stack, i)
		res += dp[i+1]
		res %= mod
	}
	return res
}

type pair struct {
	val   int
	count int
}

func sumSubarrayMins1(A []int) int {
	res, n, mod := 0, len(A), 1000000007
	lefts, rights, leftStack, rightStack := make(
		[]int,
		n,
	), make(
		[]int,
		n,
	), []*pair{}, []*pair{}
	for i := 0; i < n; i++ {
		count := 1
		for len(leftStack) != 0 && leftStack[len(leftStack)-1].val > A[i] {
			count += leftStack[len(leftStack)-1].count
			leftStack = leftStack[:len(leftStack)-1]
		}
		leftStack = append(leftStack, &pair{val: A[i], count: count})
		lefts[i] = count
	}

	for i := n - 1; i >= 0; i-- {
		count := 1
		for len(rightStack) != 0 && rightStack[len(rightStack)-1].val >= A[i] {
			count += rightStack[len(rightStack)-1].count
			rightStack = rightStack[:len(rightStack)-1]
		}
		rightStack = append(rightStack, &pair{val: A[i], count: count})
		rights[i] = count
	}

	for i := 0; i < n; i++ {
		res = (res + A[i]*lefts[i]*rights[i]) % mod
	}
	return res
}

func sumSubarrayMins2(A []int) int {
	res, mod := 0, 1000000007
	for i := 0; i < len(A); i++ {
		stack := []int{}
		stack = append(stack, A[i])
		for j := i; j < len(A); j++ {
			if stack[len(stack)-1] >= A[j] {
				stack = stack[:len(stack)-1]
				stack = append(stack, A[j])
			}
			res += stack[len(stack)-1]
		}
	}
	return res % mod
}
