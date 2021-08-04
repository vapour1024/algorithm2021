package week07

var fa []int

func findRedundantConnection(edges [][]int) []int {
	count := len(edges)
	fa = make([]int, count+1)
	for i := 1; i <= count; i++ {
		fa[i]=i
	}
	for i := 0; i < count; i++ {
		if find(fa[edges[i][0]]) != find(fa[edges[i][1]]) {
			union(fa[edges[i][0]], fa[edges[i][1]])
		} else {
			return edges[i]
		}
	}
	return nil
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
