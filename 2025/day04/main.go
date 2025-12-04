package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func canRemove(grid []string, r, c, M, N int) bool {
	if grid[r][c] != '@' {
		return false
	}
	neighborCount := 0
	for _, dr := range []int{-1, 0, 1} {
		for _, dc := range []int{-1, 0, 1} {
			if r+dr < 0 || c+dc < 0 ||
				r+dr >= M || c+dc >= N ||
				r+dr == r && c+dc == c {
				continue
			}
			if grid[r+dr][c+dc] == '@' {
				neighborCount++
			}
		}
	}
	return neighborCount < 4
}

func part1(input []string) int {
	M, N := len(input), len(input[0])
	output := 0
	for r := range input {
		for c := range input[r] {
			if canRemove(input, r, c, M, N) {
				output++
			}
		}
	}
	return output
}

func part2(input []string) int {
	M, N := len(input), len(input[0])
	output := 0
	for {
		removed := 0
		for r := range input {
			for c := range input[r] {
				if canRemove(input, r, c, M, N) {
					line := []byte(input[r])
					line[c] = '.'
					input[r] = string(line)
					removed++
				}
			}
		}
		if removed == 0 {
			break
		}
		output += removed
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
	useSample := flag.Bool("sample", false, "Use sample.txt (manually filled in)")
	flag.Parse()
	fileName := "input.txt"
	if *useSample {
		fmt.Println("USING SAMPLE INPUT FILE")
		fileName = "sample.txt"
	}
	input, err := readInput(fileName)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
