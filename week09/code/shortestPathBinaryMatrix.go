package week09

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return 1
	}
	dx, dy := []int{-1, -1, 0, 1, 1, 1, 0, -1}, []int{0, 1, 1, 1, 0, -1, -1, -1}
	queue := make([]int, 0)
	queue = append(queue, 0)
	grid[0][0] = 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		x, y := cur/n, cur%n
		for i := 0; i < 8; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 0 {
				queue = append(queue, nx*n+ny)
				grid[nx][ny] = grid[x][y] + 1
				if nx == m-1 && ny == n-1 {
					return grid[nx][ny]
				}
			}
		}
	}
	return -1
}
