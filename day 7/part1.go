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

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		target, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		numStrs := strings.Fields(parts[1])
		numbers := make([]int, len(numStrs))
		for i, str := range numStrs {
			numbers[i], _ = strconv.Atoi(str)
		}

		if canReachTarget(numbers, 1, numbers[0], target) {
			sum += target
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	fmt.Println(sum)
}

func canReachTarget(numbers []int, index, currentValue, target int) bool {

	if index == len(numbers) {
		return currentValue == target
	}

	if canReachTarget(numbers, index+1, currentValue+numbers[index], target) {
		return true
	}
	if canReachTarget(numbers, index+1, currentValue*numbers[index], target) {
		return true
	}

	return false

}
