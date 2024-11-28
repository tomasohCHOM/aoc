package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	r int
	c int
}

var dirs = map[rune][2]int{
	'^': {-1, 0},
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
}

func part1(input []string) int {
	visitedHouses := map[position]bool{{0, 0}: true}
	r, c := 0, 0
	for _, char := range input[0] {
		newDir := dirs[char]
		r, c = r+newDir[0], c+newDir[1]
		pos := position{r, c}
		_, ok := visitedHouses[pos]
		if !ok {
			visitedHouses[pos] = true
		}
	}
	return len(visitedHouses)
}

func part2(input []string) int {
	visitedHouses := map[position]bool{{0, 0}: true}
	santaR, santaC := 0, 0
	robotR, robotC := 0, 0
	for i, char := range input[0] {
		newDir := dirs[char]
		if i%2 == 0 { // Santa
			santaR, santaC = santaR+newDir[0], santaC+newDir[1]
			pos := position{santaR, santaC}
			_, ok := visitedHouses[pos]
			if !ok {
				visitedHouses[pos] = true
			}
		} else { // Robot
			robotR, robotC = robotR+newDir[0], robotC+newDir[1]
			pos := position{robotR, robotC}
			_, ok := visitedHouses[pos]
			if !ok {
				visitedHouses[pos] = true
			}
		}
	}
	return len(visitedHouses)
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
