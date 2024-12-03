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
		input += scanner.Text()
	}

	pattern := `mul\((\d+),(\d+)\)|don't\(\)|do\(\)`

	re := regexp.MustCompile(pattern)

	sum := 0
	isEnabled := true

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		if len(match) > 2 && match[0][:3] == "mul" {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			if isEnabled {
				sum += num1 * num2
			}
		}
		if match[0] == "don't()" {
			isEnabled = false
		}
		if match[0] == "do()" {
			isEnabled = true
		}
	}

	fmt.Println(sum)
}
