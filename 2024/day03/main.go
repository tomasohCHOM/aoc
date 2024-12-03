package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) int {
	output := 0
	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	matches := r.FindAllString(input, -1)
	for _, match := range matches {
		separated := match[4 : len(match)-1]
		nums := strings.Split(separated, ",")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		output += num1 * num2
	}
	return output
}

func part2(input string) int {
	output := 0
	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)|do\\(\\)|don\\'t\\(\\)")
	matches := r.FindAllString(input, -1)
	enabled := true
	for _, match := range matches {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else if enabled {
			separated := match[4 : len(match)-1]
			nums := strings.Split(separated, ",")
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			output += num1 * num2
		}
	}
	return output
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	inputAsLine := strings.Join(input, "")
	fmt.Println("Part 1:", part1(inputAsLine))
	fmt.Println("Part 2:", part2(inputAsLine))
}
