package main

import (
	"fmt"
)

func main() {
	//Valid Parenthesis String
	s := "(*)"
	fmt.Println("is valid string: ", checkValidString(s))

	//Number of Islands
	island := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	fmt.Println("Number of Islands: ", numIslands(island))

	//Minimum Path Sum
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println("Minimum path sum: ", minPathSum(grid))

	//Single Number
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println("Single Number in given Array: ", singleNumber(nums))
}
