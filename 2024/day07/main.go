package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func canMake1(i, currVal, evaluator int, nums []int) bool {
	if i == len(nums) {
		return currVal == evaluator
	}
	if currVal > evaluator {
		return false
	}
	add := canMake1(i+1, currVal+nums[i], evaluator, nums)
	multiply := canMake1(i+1, currVal*nums[i], evaluator, nums)
	return add || multiply
}

func canMake2(i, currVal, evaluator int, nums []int) bool {
	if i == len(nums) {
		return currVal == evaluator
	}
	if currVal > evaluator {
		return false
	}
	add := canMake2(i+1, currVal+nums[i], evaluator, nums)
	multiply := canMake2(i+1, currVal*nums[i], evaluator, nums)

	withConcat, _ := strconv.Atoi(strconv.Itoa(currVal) + strconv.Itoa(nums[i]))
	concatenate := canMake2(i+1, withConcat, evaluator, nums)

	return add || multiply || concatenate
}

func part1(input []string) int {
	output := 0
	for _, line := range input {
		parsedLine := strings.Split(line, ": ")
		evaluator, _ := strconv.Atoi(parsedLine[0])
		numsList := strings.Split(parsedLine[1], " ")
		nums := []int{}
		for _, elem := range numsList {
			num, _ := strconv.Atoi(elem)
			nums = append(nums, num)
		}
		if canMake1(1, nums[0], evaluator, nums) {
			output += evaluator
		}
	}
	return output
}

func part2(input []string) int {
	output := 0
	for _, line := range input {
		parsedLine := strings.Split(line, ": ")
		evaluator, _ := strconv.Atoi(parsedLine[0])
		numsList := strings.Split(parsedLine[1], " ")
		nums := []int{}
		for _, elem := range numsList {
			num, _ := strconv.Atoi(elem)
			nums = append(nums, num)
		}
		if canMake2(1, nums[0], evaluator, nums) {
			output += evaluator
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
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
