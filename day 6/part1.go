package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

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

	visited := make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[i]))
	}

	guardX, guardY := foundPositionOfGuard(grid)

	countMoves := findMoves(grid, visited, guardX, guardY)
	fmt.Println(countMoves)
}

func foundPositionOfGuard(grid [][]rune) (int, int) {
	for i, row := range grid {
		for j, cell := range row {
			if cell == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func findMoves(grid [][]rune, visited [][]int, guardX, guardY int) int {
	moves := 0

	for {
		if grid[guardX][guardY] == '^' {
			if guardX-1 >= 0 {
				if grid[guardX-1][guardY] == '.' {
					grid[guardX][guardY] = '.'
					guardX--
					grid[guardX][guardY] = '^'
					visited[guardX][guardY] = 1
				} else if grid[guardX-1][guardY] == '#' {
					grid[guardX][guardY] = '>'
				}
			} else {
				break
			}
		} else if grid[guardX][guardY] == '>' {
			if guardY+1 < len(grid) {
				if grid[guardX][guardY+1] == '.' {
					grid[guardX][guardY] = '.'
					guardY++
					grid[guardX][guardY] = '>'
					visited[guardX][guardY] = 1
				} else if grid[guardX][guardY+1] == '#' {
					grid[guardX][guardY] = 'v'
				}
			} else {
				break
			}
		} else if grid[guardX][guardY] == 'v' {
			if guardX+1 < len(grid[0]) {
				if grid[guardX+1][guardY] == '.' {
					grid[guardX][guardY] = '.'
					guardX++
					grid[guardX][guardY] = 'v'
					visited[guardX][guardY] = 1
				} else if grid[guardX+1][guardY] == '#' {
					grid[guardX][guardY] = '<'
				}
			} else {
				break
			}
		} else if grid[guardX][guardY] == '<' {
			if guardY-1 >= 0 {
				if grid[guardX][guardY-1] == '.' {
					grid[guardX][guardY] = '.'
					guardY--
					grid[guardX][guardY] = '<'
					visited[guardX][guardY] = 1
				} else if grid[guardX][guardY-1] == '#' {
					grid[guardX][guardY] = '^'
				}
			} else {
				break
			}

		}

	}

	moves = 0
	for _, row := range visited {
		for _, cell := range row {
			moves += cell
		}
	}

	return moves
}
