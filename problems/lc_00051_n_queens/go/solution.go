package lc_00051_n_queens

import "strings"

// solveNQueens returns all distinct solutions to the n-queens puzzle.
// Each solution represents a distinct board configuration where 'Q' indicates a queen and '.' indicates an empty space.
func solveNQueens(n int) [][]string {
	// Implementation will be added by you
	s := make([]string, n)
	for i := 0; i < n; i += 1 {
		s[i] = "."
	}

	one := make([]string, 0)
	all := make([][]string, 0)
	used := make([][]bool, n)

	for i := 0; i < n; i += 1 {
		used[i] = make([]bool, n)
	}

	back(n, &one, &all, used, s)
	return all
}

func back(n int, one *[]string, all *[][]string, used [][]bool, s []string) {
	if n == 0 {
		return
	}

	if len(*one) == n {
		cp := make([]string, len(*one))
		copy(cp, *one)
		*all = append(*all, cp)
		return
	}

	size := len(*one)
	for i := 0; i < n; i += 1 {
		if isAttack(used, i, size, n) {
			continue
		}

		used[i][size] = true
		s[i] = "Q"
		*one = append(*one, strings.Join(s, ""))
		s[i] = "."
		back(n, one, all, used, s)
		used[i][size] = false
		*one = (*one)[:len(*one)-1]
	}
}

func isAttack(used [][]bool, x, y, n int) bool {
	if x < 0 || x >= n || y < 0 || y >= n {
		return false
	}
	// 判断是否可能被攻击
	for i := y - 1; i >= 0; i -= 1 {
		if used[x][i] {
			return true
		}
	}

	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if used[i][j] {
			return true
		}
	}

	for i, j := x+1, y-1; i < n && j >= 0; i, j = i+1, j-1 {
		if used[i][j] {
			return true
		}
	}

	return false
}
