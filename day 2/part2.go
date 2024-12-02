package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafe(report []int) bool {

	isIncreasing := true
	isDecreasing := true

	for i := 0; i < len(report)-1; i++ {

		diff := report[i+1] - report[i]
		temp_diff := diff
		diff = abs(diff)

		if diff > 3 || diff == 0 {
			return false
		}

		if temp_diff < 0 {
			isIncreasing = false
		}
		if temp_diff > 0 {
			isDecreasing = false
		}
	}

	isSafe := isIncreasing || isDecreasing
	return isSafe
}
func generateReports(report []int) [][]int {

	var reports [][]int

	for i := 0; i < len(report); i++ {

		newReport := append([]int(nil), report[:i]...)
		newReport = append(newReport, report[i+1:]...)

		reports = append(reports, newReport)
	}
	return reports
}

func countSafeReports(data [][]int) int {
	count := 0

	for _, report := range data {

		if isSafe(report) {
			count++
			continue
		}

		reports := generateReports(report)
		for _, newReport := range reports {
			if isSafe(newReport) {
				count++
				break
			}
		}
	}

	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var data [][]int

	for {
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}

		strNumbers := strings.Fields(line)
		var report []int
		for _, str := range strNumbers {
			num, err := strconv.Atoi(str)
			if err != nil {
				return
			}
			report = append(report, num)
		}
		data = append(data, report)
	}

	safeCount := countSafeReports(data)
	fmt.Println(safeCount)
}
