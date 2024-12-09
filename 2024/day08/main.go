package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type coords struct {
	x, y int
}

func part1(input []string) int {
	M, N := len(input), len(input[0])
	antennaGroups := map[rune][]coords{}
	for r, line := range input {
		for c, char := range line {
			if char != '.' {
				antennaGroups[char] = append(antennaGroups[char], coords{r, c})
			}
		}
	}
	locations := map[coords]bool{}
	for _, coordinates := range antennaGroups {
		sort.Slice(coordinates, func(i, j int) bool {
			return coordinates[i].x < coordinates[j].x
		})
		for i := 0; i < len(coordinates); i++ {
			for j := i + 1; j < len(coordinates); j++ {
				xDiff := coordinates[j].x - coordinates[i].x
				yDiff := coordinates[j].y - coordinates[i].y
				if coordinates[i].x-xDiff >= 0 && coordinates[i].y-yDiff >= 0 && coordinates[i].y-yDiff < M {
					locations[coords{coordinates[i].x - xDiff, coordinates[i].y - yDiff}] = true
				}
				if coordinates[j].x+xDiff < N && coordinates[j].y+yDiff >= 0 && coordinates[j].y+yDiff < M {
					locations[coords{coordinates[j].x + xDiff, coordinates[j].y + yDiff}] = true
				}
			}
		}
	}
	return len(locations)
}

func part2(input []string) int {
	M, N := len(input), len(input[0])
	antennaGroups := map[rune][]coords{}
	for r, line := range input {
		for c, char := range line {
			if char != '.' {
				antennaGroups[char] = append(antennaGroups[char], coords{r, c})
			}
		}
	}
	locations := map[coords]bool{}
	for _, coordinates := range antennaGroups {
		sort.Slice(coordinates, func(i, j int) bool {
			return coordinates[i].x < coordinates[j].x
		})
		for i := 0; i < len(coordinates); i++ {
			for j := i + 1; j < len(coordinates); j++ {
				xDiff := coordinates[j].x - coordinates[i].x
				yDiff := coordinates[j].y - coordinates[i].y
				currCoords := coords{coordinates[i].x, coordinates[i].y}

				for currCoords.x-xDiff >= 0 && currCoords.y-yDiff >= 0 && currCoords.y-yDiff < M {
					currCoords.x -= xDiff
					currCoords.y -= yDiff
				}
				for currCoords.x < N && currCoords.y >= 0 && currCoords.y < M {
					locations[coords{currCoords.x, currCoords.y}] = true
					currCoords.x += xDiff
					currCoords.y += yDiff
				}
			}
		}
	}
	return len(locations)
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
