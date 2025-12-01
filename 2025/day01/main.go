package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var rotation map[byte]int = map[byte]int{'L': -1, 'R': 1}

func part1(input []string) int {
	pointing, output := 50, 0
	for _, line := range input {
		delta := rotation[line[0]]
		num, _ := strconv.Atoi(line[1:])
		pointing += delta * (num % 100)
		pointing = (pointing + 100) % 100
		if pointing == 0 {
			output++
		}
	}
	return output
}

func part2(input []string) int {
	pointing, output := 50, 0
	for _, line := range input {
		delta := rotation[line[0]]
		num, _ := strconv.Atoi(line[1:])
		prev := pointing
		pointing += delta * num
		crossing := 0
		if pointing <= 0 && prev != 0 {
			crossing = 1
		}
		output += crossing + abs(pointing)/100
		pointing = ((pointing % 100) + 100) % 100
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
