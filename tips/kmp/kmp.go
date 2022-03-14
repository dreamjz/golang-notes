package kmp

// BruteForce 暴力匹配
func bruteForce(s string, p string) int {
	m, n := len(s), len(p)
	i, j := 0, 0
	for i < m && j < n {
		if s[i] == p[j] {
			i++
			j++
		} else {
			i = i - j + 1 // 回到起始点的下一个
			j = 0
		}
	}
	if j == n {
		return i - j // 返回匹配的起始位置
	}
	return -1
}

// kmp KMP 算法
func kmp(s string, t string) int {
	next := buildNext(t)
	m, n := len(s), len(t)
	i, j := 0, 0
	for i < m && j < n {
		if j == -1 || s[i] == t[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == n {
		return i - j
	}
	return -1
}

func buildNext(t string) []int {
	n := len(t)
	j, k := 0, -1
	next := make([]int, len(t))
	next[j] = k
	for j < n-1 {
		if k == -1 || t[k] == t[j] {
			j++
			k++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next
}
