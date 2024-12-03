package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	pattern := `mul\((\d+),(\d+)\)`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)

	sum := 0

	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}

	fmt.Println(sum)
}
