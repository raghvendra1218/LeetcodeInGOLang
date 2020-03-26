package main

import "math"

func minPathSum(grid [][]int) int {
	row, col := 0, 0
	minWeight := math.MaxInt64

	//define a memo matrix of grid size
	n, m := len(grid), len(grid[0])
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, m)
	}

	//Initialize teh memo matrix with -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			memo[i][j] = -1
		}
	}

	//recurse through grid to get the minWeight value
	minWeight = recurse(row, col, memo, grid)

	return minWeight
}

//recurse function for calculating the min-weight in one path
func recurse(row, col int, memo, grid [][]int) int {
	//Base case

	//1. If row and col reach out of bound of grid, it is safe to return the Math.MaxInt64 value because in that case
	//	 we will go to the inbound direction (try with diagram)
	if row == len(grid) || col == len(grid[0]) {
		return math.MaxInt64
	}

	//2. If row and col reaches at the last point in the matrix, then the minPath for that point is value at that row and col
	if row == len(grid)-1 && col == len(grid[0])-1 {
		return grid[row][col]
	}

	if memo[row][col] == -1 {
		memo[row][col] = grid[row][col] + min(recurse(row+1, col, memo, grid), recurse(row, col+1, memo, grid))
	}

	return memo[row][col]
}

//min function for calculating minimum between two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
