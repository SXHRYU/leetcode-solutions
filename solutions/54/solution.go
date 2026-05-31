package leetcode

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	if len(matrix) == 1 {
		for i := 0; i < len(matrix[0]); i++ {
			res = append(res, matrix[0][i])
		}
		return res
	}
	if len(matrix[0]) == 1 {
		for i := 0; i < len(matrix); i++ {
			res = append(res, matrix[i][0])
		}
		return res
	}
	visit, m, n, round, x, y, spDir := make(
		[][]int,
		len(matrix),
	), len(
		matrix,
	), len(
		matrix[0],
	), 0, 0, 0, [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	for i := 0; i < m; i++ {
		visit[i] = make([]int, n)
	}
	visit[x][y] = 1
	res = append(res, matrix[x][y])
	for i := 0; i < m*n; i++ {
		x += spDir[round%4][0]
		y += spDir[round%4][1]
		if (x == 0 && y == n-1) || (x == m-1 && y == n-1) || (y == 0 && x == m-1) {
			round++
		}
		if visit[x][y] == 0 {
			visit[x][y] = 1
			res = append(res, matrix[x][y])
		}
		switch round % 4 {
		case 0:
			if y+1 <= n-1 && visit[x][y+1] == 1 {
				round++
				continue
			}
		case 1:
			if x+1 <= m-1 && visit[x+1][y] == 1 {
				round++
				continue
			}
		case 2:
			if y-1 >= 0 && visit[x][y-1] == 1 {
				round++
				continue
			}
		case 3:
			if x-1 >= 0 && visit[x-1][y] == 1 {
				round++
				continue
			}
		}
	}
	return res
}

func spiralOrder2(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}

	n := len(matrix[0])
	if n == 0 {
		return nil
	}

	top, left, bottom, right := 0, 0, m-1, n-1
	count, sum := 0, m*n
	res := []int{}

	for count < sum {
		i, j := top, left
		for j <= right && count < sum {
			res = append(res, matrix[i][j])
			count++
			j++
		}
		i, j = top+1, right
		for i <= bottom && count < sum {
			res = append(res, matrix[i][j])
			count++
			i++
		}
		i, j = bottom, right-1
		for j >= left && count < sum {
			res = append(res, matrix[i][j])
			count++
			j--
		}
		i, j = bottom-1, left
		for i > top && count < sum {
			res = append(res, matrix[i][j])
			count++
			i--
		}

		top, left, bottom, right = top+1, left+1, bottom-1, right-1
	}
	return res
}
