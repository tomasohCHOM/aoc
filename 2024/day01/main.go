package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(input []string) int {
	firstColNums := []int{}
	secondColNums := []int{}
	for _, line := range input {
		nums := strings.Fields(line)
		firstColNum, _ := strconv.Atoi(nums[0])
		secondColNum, _ := strconv.Atoi(nums[1])

		firstColNums = append(firstColNums, firstColNum)
		secondColNums = append(secondColNums, secondColNum)
	}
	slices.Sort(firstColNums)
	slices.Sort(secondColNums)
	output := 0
	for i := range firstColNums {
		output += int(math.Abs(float64(firstColNums[i]) - float64(secondColNums[i])))
	}
	return output
}

func part2(input []string) int {
	freqs := map[int]int{}
	for _, line := range input {
		nums := strings.Fields(line)
		secondColNum, _ := strconv.Atoi(nums[1])
		_, ok := freqs[secondColNum]
		if !ok {
			freqs[secondColNum] = 1
		} else {
			freqs[secondColNum] += 1
		}
	}
	output := 0
	for _, line := range input {
		nums := strings.Fields(line)
		firstColNum, _ := strconv.Atoi(nums[0])
		freq, ok := freqs[firstColNum]
		if ok {
			output += firstColNum * freq
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
