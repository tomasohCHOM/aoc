package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func part1(input any) int {
	return 0
}

func part2(input any) int {
	return 0
}

func parseInput(filename string) (any, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var parsed []string
	scanner := bufio.NewScanner(file)

	// Use any parsing method here
	for scanner.Scan() {
		parsed = append(parsed, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return parsed, nil
}

func main() {
	useSample := flag.Bool("sample", false, "Use sample.txt (manually filled in)")
	flag.Parse()
	fileName := "input.txt"
	if *useSample {
		fmt.Println("USING SAMPLE INPUT FILE")
		fileName = "sample.txt"
	}
	input, err := parseInput(fileName)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
