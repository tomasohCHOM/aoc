package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getStones(input string) []int {
	parsedLine := strings.Split(input, " ")
	stones := []int{}
	for _, elem := range parsedLine {
		val, _ := strconv.Atoi(elem)
		stones = append(stones, val)
	}
	return stones
}

func part1(input string) int {
	stones := getStones(input)
	for i := 0; i < 25; i++ {
		newStones := []int{}
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if stoneStr := strconv.Itoa(stone); len(stoneStr)%2 == 0 {
				left, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
				right, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
				newStones = append(newStones, left, right)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}
	return len(stones)
}

func part2(input string) int {
	parsedLine := strings.Split(input, " ")
	stones := map[int]int{}
	for _, elem := range parsedLine {
		val, _ := strconv.Atoi(elem)
		stones[val]++
	}
	for i := 0; i < 75; i++ {
		newStones := map[int]int{}
		for stone, count := range stones {
			if stone == 0 {
				newStones[1] += count
			} else if stoneStr := strconv.Itoa(stone); len(stoneStr)%2 == 0 {
				left, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
				right, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
				newStones[left] += count
				newStones[right] += count
			} else {
				newStones[stone*2024] += count
			}
		}
		stones = newStones
	}
	output := 0
	for _, v := range stones {
		output += v
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
	fmt.Println("Part 1:", part1(input[0]))
	fmt.Println("Part 2:", part2(input[0]))
}
