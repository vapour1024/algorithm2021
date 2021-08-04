package week07

import "fmt"

var fa []int

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	mark := make([][]int, m)
	for i := 0; i < m; i++ {
		mark[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mark[i][j] = i*n + j
		}
	}
	fmt.Println(mark)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				mark[i][j] = 0
			}
		}
	}
	return 0
}
func union(i, j int) {
	x := find(i)
	y := find(j)
	if x != y {
		fa[x] = y
	}
}
func find(x int) int {
	if x == fa[x] {
		return x
	}
	fa[x] = find(fa[x])
	return fa[x]
}
