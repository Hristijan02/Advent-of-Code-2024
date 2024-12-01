package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

	var leftNums []int
	var rightNums []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		left, _ := strconv.Atoi(numbers[0])
		righ, _ := strconv.Atoi(numbers[1])

		leftNums = append(leftNums, left)
		rightNums = append(rightNums, righ)
	}

	sort.Slice(leftNums, func(i, j int) bool {
		return leftNums[i] < leftNums[j]
	})

	sort.Slice(rightNums, func(i, j int) bool {
		return rightNums[i] < rightNums[j]
	})

	var sum int

	for i := 0; i < len(leftNums); i++ {
		temp_sum := abs(leftNums[i] - rightNums[i])
		sum += temp_sum
	}

	fmt.Println(sum)

}
