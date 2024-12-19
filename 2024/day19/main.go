package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Word Break Problem
func patternCount(s string, patternSet map[string]bool) int {
	N := len(s)
	dp := make([]int, N+1)
	dp[0] = 1

	for i := 1; i < N+1; i++ {
		for j := 0; j < i; j++ {
			if patternSet[s[j:i]] {
				dp[i] += dp[j]
			}
		}
	}
	return dp[N]
}

func solveParts(input []string) (int, int) {
	patterns := strings.Split(input[0], ", ")
	patternSet := map[string]bool{}
	for _, pattern := range patterns {
		patternSet[pattern] = true
	}
	possible, totalNumWays := 0, 0
	for i := 2; i < len(input); i++ {
		numWays := patternCount(input[i], patternSet)
		if numWays > 0 {
			possible++
			totalNumWays += numWays
		}
	}
	return possible, totalNumWays
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
	possible, totalNumWays := solveParts(input)
	fmt.Println("Part 1:", possible)
	fmt.Println("Part 2:", totalNumWays)
}
