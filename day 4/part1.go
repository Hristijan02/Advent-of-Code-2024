package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var directions = [8][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

func countWords(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	wordLength := len(word)
	counter := 0
	wordRunes := []rune(word)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				isMatch := true
				for i := 0; i < wordLength; i++ {
					newR := r + dir[0]*i
					newC := c + dir[1]*i
					if newR < 0 || newR >= rows || newC < 0 || newC >= cols || grid[newR][newC] != wordRunes[i] {
						isMatch = false
						break
					}
				}
				if isMatch {
					counter++
				}
			}
		}
	}
	return counter
}

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

	word := "XMAS"

	result := countWords(grid, word)

	fmt.Println(result)
}
