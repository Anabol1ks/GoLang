package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Edge struct {
	u, v, w int
}

func find(parent []int, x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent, parent[x])
	return parent[x]
}

func union(parent []int, rank []int, x, y int) bool {
	rootX := find(parent, x)
	rootY := find(parent, y)

	if rootX != rootY {
		if rank[rootX] > rank[rootY] {
			parent[rootY] = rootX
		} else if rank[rootX] < rank[rootY] {
			parent[rootX] = rootY
		} else {
			parent[rootY] = rootX
			rank[rootX]++
		}
		return true
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(in, &n)
		c := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &c[j])
		}
		fmt.Fscan(in, &m)

		edges := make([]Edge, 0, m+n)

		for j := 0; j < n; j++ {
			edges = append(edges, Edge{0, j + 1, c[j]})
		}

		for j := 0; j < m; j++ {
			var u, v, w int
			fmt.Fscan(in, &u, &v, &w)
			edges = append(edges, Edge{u, v, w})
		}

		// Сортируем рёбра по весу
		sort.Slice(edges, func(i, j int) bool {
			return edges[i].w < edges[j].w
		})

		parent := make([]int, n+1)
		rank := make([]int, n+1)
		for j := 0; j <= n; j++ {
			parent[j] = j
		}

		totalCost := 0
		for _, edge := range edges {
			if union(parent, rank, edge.u, edge.v) {
				totalCost += edge.w
			}
		}

		fmt.Fprintln(out, totalCost)
	}
}
