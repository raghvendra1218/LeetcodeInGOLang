package main

func numIslands(grid [][]byte) int {
	numberOfIslands := 0
	//traverse through grid and we will add to the numberOfIslands as we find an island
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '1' {
				traversegrid(grid, row, col)
				numberOfIslands++
			}
		}
	}
	return numberOfIslands
}

func traversegrid(grid [][]byte, row, col int) {

	//Boundary conditions or the case when current grid's position is zero
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == '0' {
		return
	}

	//Mark the current grid's position as visited, we can change it to 0, so that it is not counted again while traversal
	grid[row][col] = '0'
	//Do a Depth First Search
	traversegrid(grid, row, col+1)
	traversegrid(grid, row, col-1)
	traversegrid(grid, row+1, col)
	traversegrid(grid, row-1, col)
}
