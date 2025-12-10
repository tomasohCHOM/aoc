package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func getCoords(input []string) [][2]int {
	coords := [][2]int{}
	for _, line := range input {
		coord := strings.Split(line, ",")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		coords = append(coords, [2]int{x, y})
	}
	return coords
}

func part1(input []string) int {
	coords := getCoords(input)
	output := 0
	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			x1, y1 := coords[i][0], coords[i][1]
			x2, y2 := coords[j][0], coords[j][1]
			output = max(output, (abs(x2-x1)+1)*(abs(y2-y1)+1))
		}
	}
	return output
}

func part2(input []string) int {
	coords := getCoords(input)
	sortedCoords := make([][2]int, len(coords))
	copy(sortedCoords, coords)

	sort.Slice(sortedCoords, func(i, j int) bool {
		if sortedCoords[i][0] != sortedCoords[j][0] {
			return sortedCoords[i][0] < sortedCoords[j][0]
		}
		return sortedCoords[i][1] < sortedCoords[j][1]
	})

	seenX, seenY := map[int]int{}, map[int]int{}
	rankX, rankY := 0, 0
	compacted := [][2]int{}
	for _, coord := range sortedCoords {
		x, y := coord[0], coord[1]
		if _, ok := seenX[x]; !ok {
			rankX++
		}
		if _, ok := seenY[y]; !ok {
			rankY++
		}
		seenX[x] = rankX
		seenY[y] = rankY
		compacted = append(compacted, [2]int{rankX, rankY})
	}

	M, N := rankY+1, rankX+1
	grid := make([][]byte, M)
	for i := range M {
		row := make([]byte, N)
		for j := range N {
			row[j] = '.'
		}
		grid[i] = row
	}
	// --- draw points + connecting segments using original coords ---
	for i := range coords {
		x1, y1 := coords[i][0], coords[i][1]
		cx1 := seenX[x1]
		cy1 := seenY[y1]

		// place original point
		grid[cy1][cx1] = '#'

		if i > 0 {
			// previous point
			x0, y0 := coords[i-1][0], coords[i-1][1]

			// ONLY connect if perfectly horizontal or vertical
			if x0 == x1 || y0 == y1 {

				cx0 := seenX[x0]
				cy0 := seenY[y0]

				// horizontal line
				if y0 == y1 {
					if cx0 <= cx1 {
						for c := cx0; c <= cx1; c++ {
							grid[cy0][c] = '#'
						}
					} else {
						for c := cx1; c <= cx0; c++ {
							grid[cy0][c] = '#'
						}
					}
				}

				// vertical line
				if x0 == x1 {
					if cy0 <= cy1 {
						for r := cy0; r <= cy1; r++ {
							grid[r][cx0] = '#'
						}
					} else {
						for r := cy1; r <= cy0; r++ {
							grid[r][cx0] = '#'
						}
					}
				}
			}
		}
	}

	printGrid(grid)

	output := 0
	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			x1, y1 := coords[i][0], coords[i][1]
			x2, y2 := coords[j][0], coords[j][1]
			output = max(output, (abs(x2-x1)+1)*(abs(y2-y1)+1))
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
