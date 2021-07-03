package main

func numIslands(grid [][]byte) int {
	count := 0
	for i, v := range grid {
		for j, vv := range v {
			//printGrid(grid)
			if vv == '0' || vv == '9' {
				continue
			}
			dfs(i, j, &grid)
			count++
		}
	}
	return count
}

func dfs(row, col int, grid *[][]byte) {
	if row < 0 || row >= len(*grid) || col < 0 || col >= len((*grid)[0]) {
		return
	}

	if (*grid)[row][col] != '1' {
		return
	}

	(*grid)[row][col] = '9' // mark as visited
	dfs(row+1, col, grid)
	dfs(row-1, col, grid)
	dfs(row, col+1, grid)
	dfs(row, col-1, grid)
}

func printGrid(grid [][]byte) {
	for _, x := range grid {
		fmt.Printf("[%s]\n", string(x))
	}
	fmt.Println()
}
