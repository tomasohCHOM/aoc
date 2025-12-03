package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func findIndexOfHighest(bank string, start, end int) int {
	highest := start
	for i := highest; i < end; i++ {
		if bank[i] > bank[highest] {
			highest = i
		}
		if bank[highest] == '9' {
			return highest
		}
	}
	return highest
}

func part1(input []string) int {
	output := 0
	for _, bank := range input {
		firstIdx := findIndexOfHighest(bank, 0, len(bank)-1)
		secondIdx := findIndexOfHighest(bank, firstIdx+1, len(bank))
		output += int(bank[firstIdx]-'0')*10 + int(bank[secondIdx]-'0')
	}
	return output
}

func part2(input []string) int {
	output := 0
	for _, bank := range input {
		digitIdx, joltage := 0, 0
		for i := 11; i >= 0; i-- {
			digitIdx = findIndexOfHighest(bank, digitIdx, len(bank)-i)
			digit := int(bank[digitIdx] - '0')
			joltage = joltage*10 + digit
			digitIdx++
		}
		output += joltage
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
