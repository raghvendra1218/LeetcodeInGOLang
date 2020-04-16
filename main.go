package main

import (
	"fmt"
)

func main() {
	s := "(*)"
	fmt.Println("is valid string: ", checkValidString(s))
	//time.Sleep(1 * time.Second)
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	fmt.Println("Number of Islands: ", numIslands(grid))
}
