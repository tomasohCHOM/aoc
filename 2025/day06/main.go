package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(input []string) int {
	var rows [][]int
	for r := 0; r < len(input)-1; r++ {
		fields := strings.Fields(input[r])
		nums := make([]int, len(fields))
		for i, f := range fields {
			nums[i], _ = strconv.Atoi(f)
		}
		rows = append(rows, nums)
	}
	ops := strings.Fields(input[len(input)-1])
	output := 0
	for c, op := range ops {
		value := 0
		if op == "*" {
			value = 1
		}
		for r := range rows {
			if op == "+" {
				value += rows[r][c]
			} else {
				value *= rows[r][c]
			}
		}
		output += value
	}
	return output
}

func part2(input []string) int {
	M, N := len(input), len(input[0])
	opRow := input[M-1]
	var opRanges [][2]int
	start := 0
	for i := 1; i < N; i++ {
		if opRow[i] != ' ' {
			opRanges = append(opRanges, [2]int{start, i - 1})
			start = i
		}
	}
	opRanges = append(opRanges, [2]int{start, N})

	readNumber := func(c int) int {
		num := 0
		for r := 0; r < M-1; r++ {
			if input[r][c] != ' ' {
				digit := int(input[r][c] - '0')
				num = num*10 + digit
			}
		}
		return num
	}

	output := 0
	for _, rg := range opRanges {
		lower, upper := rg[0], rg[1]
		op := opRow[lower]
		value := 0
		if op == '*' {
			value = 1
		}
		for c := lower; c < upper; c++ {
			num := readNumber(c)
			if op == '+' {
				value += num
			} else {
				value *= num
			}
		}
		output += value
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
