package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func part1(input []string) int {
	M, N := len(input), len(input[0])
	prevRow, sPos := []byte(input[0]), strings.IndexByte(input[0], 'S')
	prevRow[sPos] = '|'
	output := 0
	for r := 1; r < M; r++ {
		currRow := prevRow
		for c, ch := range input[r] {
			if ch != '^' {
				continue
			}
			if prevRow[c] == '|' {
				output++
				currRow[c] = '.'
				if c-1 >= 0 {
					currRow[c-1] = '|'
				}
				if c+1 < N {
					currRow[c+1] = '|'
				}
			}
		}
		prevRow = currRow
	}
	return output
}

type Cell struct {
	Row int
	Col int
}

func part2(input []string) int {
	M, N := len(input), len(input[0])
	var dp map[Cell]int

	var dfs func(cell Cell) int
	dfs = func(cell Cell) int {
		if cell.Row == M {
			return 1
		}
		if cell.Col < 0 || cell.Col >= N {
			return 0
		}
		if val, ok := dp[cell]; ok {
			return val
		}
		output := 0
		if input[cell.Row][cell.Col] == '^' {
			output += dfs(Cell{Row: cell.Row + 1, Col: cell.Col - 1})
			output += dfs(Cell{Row: cell.Row + 1, Col: cell.Col + 1})
		} else {
			output += dfs(Cell{Row: cell.Row + 1, Col: cell.Col})
		}
		dp[cell] = output
		return dp[cell]
	}

	cell := Cell{Row: 0, Col: strings.IndexByte(input[0], 'S')}
	dp = make(map[Cell]int)
	output := dfs(cell)
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
