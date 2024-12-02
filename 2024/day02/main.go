package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isSafe(nums []int) bool {
	increasing, decreasing := false, false
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return false
		}
		diff := int(math.Abs(float64(nums[i]) - float64(nums[i-1])))
		if diff < 1 || diff > 3 {
			return false
		}
		if nums[i] < nums[i-1] {
			if increasing {
				return false
			}
			decreasing = true
		} else {
			if decreasing {
				return false
			}
			increasing = true
		}
	}
	return true
}

func part1(input []string) int {
	output := 0
	for _, line := range input {
		nums := []int{}
		numLine := strings.Split(line, " ")
		for _, v := range numLine {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
		if isSafe(nums) {
			output += 1
		}
	}
	return output
}

func part2(input []string) int {
	output := 0
	for _, line := range input {
		nums := []int{}
		numLine := strings.Split(line, " ")
		for _, v := range numLine {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
		for i := 0; i < len(nums); i++ {
			numsCopy := make([]int, 0, len(nums)-1)
			numsCopy = append(numsCopy, nums[:i]...)
			numsCopy = append(numsCopy, nums[i+1:]...)

			if isSafe(numsCopy) {
				output += 1
				break
			}
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
