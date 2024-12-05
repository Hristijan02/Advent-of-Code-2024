package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	dictionary := make(map[int][]int)

	var groups [][]int

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				continue
			}

			x, err1 := strconv.Atoi(parts[0])
			y, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				dictionary[x] = append(dictionary[x], y)
			}
		} else if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			var group []int
			for _, part := range parts {
				num, err := strconv.Atoi(strings.TrimSpace(part))
				if err == nil {
					group = append(group, num)
				}
			}
			groups = append(groups, group)
		}
	}

	fmt.Println(calculateMiddleSum(dictionary, groups))
}

func calculateMiddleSum(dictionary map[int][]int, groups [][]int) int {

	sum := 0
	for _, group := range groups {
		isOk := true
		for i, num := range group {
			for j := 0; j < i; j++ {
				if values, exists := dictionary[group[j]]; exists {
					if !contains(values, num) {
						isOk = false
						break
					}
				} else {
					isOk = false
					break
				}
			}
			if !isOk {
				break
			}
		}
		if isOk {
			median := group[len(group)/2]
			sum += median
		}
	}
	return sum
}

func contains(values []int, num int) bool {
	for _, value := range values {
		if value == num {
			return true
		}
	}
	return false
}
