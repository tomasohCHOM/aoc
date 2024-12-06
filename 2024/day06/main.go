package main

import (
	"bufio"
	"fmt"
	"os"
)

type coords struct {
	x int
	y int
}

func findStartingPostion(input []string) (int, int) {
	for r, line := range input {
		for c, cell := range line {
			if cell == '^' {
				return r, c
			}
		}
	}
	panic("No starting position found")
}

// Directions when rotating 90 degrees
var directions [][]int = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func part1(input []string) int {
	M, N := len(input), len(input[0])
	i, j := findStartingPostion(input)
	currDirIdx := 0

	visited := map[coords]bool{}
	visited[coords{i, j}] = true

	for {
		nextX, nextY := i+directions[currDirIdx][0], j+directions[currDirIdx][1]
		if nextX < 0 || nextY < 0 || nextX == M || nextY == N {
			break
		}
		if input[nextX][nextY] == '#' {
			currDirIdx = (currDirIdx + 1) % 4
			continue
		}
		visited[coords{nextX, nextY}] = true
		i += directions[currDirIdx][0]
		j += directions[currDirIdx][1]
	}
	return len(visited)
}

func part2(input []string) int {
	M, N := len(input), len(input[0])
	i, j := findStartingPostion(input)
	currDirIdx := 0

	possible := map[coords]bool{}
	checked := map[coords]bool{}

	isValid := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < M && y < N
	}

	containsCycle := func(x, y int) bool {
		slowX, slowY := findStartingPostion(input)
		fastX, fastY := findStartingPostion(input)
		slowDirIdx, fastDirIdx := 0, 0

		for {
			// Handle fast pointer first
			for step := 0; step < 2; step++ {
				nextFastX, nextFastY := fastX+directions[fastDirIdx][0], fastY+directions[fastDirIdx][1]
				if !isValid(nextFastX, nextFastY) {
					return false
				}
				if input[nextFastX][nextFastY] == '#' || (nextFastX == x && nextFastY == y) {
					fastDirIdx = (fastDirIdx + 1) % 4
					continue
				}
				fastX, fastY = nextFastX, nextFastY
			}

			// Handle slow pointer
			nextSlowX, nextSlowY := slowX+directions[slowDirIdx][0], slowY+directions[slowDirIdx][1]
			if !isValid(nextSlowX, nextSlowY) {
				return false
			}
			if input[nextSlowX][nextSlowY] == '#' || (nextSlowX == x && nextSlowY == y) {
				slowDirIdx = (slowDirIdx + 1) % 4
			} else {
				slowX, slowY = nextSlowX, nextSlowY
			}

			// If they meet and have the same direction
			if slowX == fastX && slowY == fastY && slowDirIdx == fastDirIdx {
				return true
			}
		}
	}

	for {
		nextX, nextY := i+directions[currDirIdx][0], j+directions[currDirIdx][1]
		if !isValid(nextX, nextY) {
			break
		}
		if input[nextX][nextY] == '#' {
			currDirIdx = (currDirIdx + 1) % 4
			continue
		}
		pos := coords{nextX, nextY}
		_, ok := checked[pos]
		if !ok && containsCycle(nextX, nextY) {
			possible[pos] = true
		}
		checked[pos] = true
		i = nextX
		j = nextY
	}
	return len(possible)
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
