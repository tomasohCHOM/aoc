package main

import (
	"bufio"
	"fmt"
	"os"
)

type tuple struct {
	first  int
	second int
}

func part1(input []string) int {
	M, N := len(input), len(input[0])
	output, word := 0, "XMAS"
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	var dfs func(i, r, c, dr, dc int)
	dfs = func(i, r, c, dr, dc int) {
		if i == len(word) {
			output += 1
			return
		}
		if r < 0 || c < 0 || r >= M || c >= N || input[r][c] != word[i] {
			return
		}
		dfs(i+1, r+dr, c+dc, dr, dc)
	}

	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			for _, dir := range directions {
				dr, dc := dir[0], dir[1]
				dfs(0, r, c, dr, dc)
			}
		}
	}
	return output
}

func part2(input []string) int {
	M, N := len(input), len(input[0])
	output := 0
	for r := 0; r < M-2; r++ {
		for c := 0; c < N-2; c++ {
			// Brute force this shit
			if input[r+1][c+1] == 'A' &&
				((input[r][c] == 'M' && input[r+2][c+2] == 'S' && input[r+2][c] == 'M' && input[r][c+2] == 'S') ||
					(input[r][c] == 'M' && input[r+2][c+2] == 'S' && input[r+2][c] == 'S' && input[r][c+2] == 'M') ||
					(input[r][c] == 'S' && input[r+2][c+2] == 'M' && input[r+2][c] == 'M' && input[r][c+2] == 'S') ||
					(input[r][c] == 'S' && input[r+2][c+2] == 'M' && input[r+2][c] == 'S' && input[r][c+2] == 'M')) {
				output += 1
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
