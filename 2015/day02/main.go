package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(input []string) int {
	output := 0
	for _, line := range input {
		dimensions := strings.Split(line, "x")
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])
		output += 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
	}
	return output
}

func part2(input []string) int {
	output := 0
	for _, line := range input {
		dimensions := strings.Split(line, "x")
		x, _ := strconv.Atoi(dimensions[0])
		y, _ := strconv.Atoi(dimensions[1])
		z, _ := strconv.Atoi(dimensions[2])
		output += min(2*x+2*y, 2*x+2*z, 2*y+2*z) + x*y*z
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
