package main

import (
	"bufio"
	"fmt"
	"os"
)

func exist(grid [][]byte, word string) bool {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false
	}

	rows := len(grid)
	cols := len(grid[0])

	if len(word) > rows*cols {
		return false
	}

	for i := 0; i < rows; i++ {

		for j := 0; j < cols; j++ {

			if grid[i][j] == word[0] {

				if dfs(grid, word, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(grid [][]byte, word string, i, j, index int) bool {

	if index == len(word) {
		return true
	}

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != word[index] {
		return false
	}

	temp := grid[i][j]
	grid[i][j] = '#'

	found := dfs(grid, word, i+1, j, index+1) ||
		dfs(grid, word, i-1, j, index+1) ||
		dfs(grid, word, i, j+1, index+1) ||
		dfs(grid, word, i, j-1, index+1)

	grid[i][j] = temp
	return found
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
