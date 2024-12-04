package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWordsX(grid [][]rune) int {

	counter := 0
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			if grid[r][c] == 'A' && grid[r-1][c-1] == 'S' && grid[r+1][c+1] == 'M' ||
				grid[r][c] == 'A' && grid[r-1][c-1] == 'M' && grid[r+1][c+1] == 'S' {

				if grid[r-1][c+1] == 'S' && grid[r+1][c-1] == 'M' ||
					grid[r-1][c+1] == 'M' && grid[r+1][c-1] == 'S' {
					counter++
				}
			}
		}
	}
	return counter
}

func main() {
	reader := bufio.NewReader((os.Stdin))

	var grid [][]rune
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		grid = append(grid, []rune(line))
	}

	result := countWordsX(grid)

	fmt.Println(result)
}
