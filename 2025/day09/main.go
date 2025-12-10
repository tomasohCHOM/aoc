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

func rectArea(x1, y1, x2, y2 int) int {
	return (abs(x2-x1) + 1) * (abs(y2-y1) + 1)
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
			output = max(output, rectArea(x1, y1, x2, y2))
		}
	}
	return output
}

func unique(a []int) []int {
	if len(a) == 0 {
		return a
	}
	out := []int{a[0]}
	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			out = append(out, a[i])
		}
	}
	return out
}

func compress(coords [][2]int) (cx, cy map[int]int, xs, ys []int) {
	xs = make([]int, len(coords))
	ys = make([]int, len(coords))

	for i, p := range coords {
		xs[i] = p[0]
		ys[i] = p[1]
	}

	sort.Ints(xs)
	sort.Ints(ys)
	xs = unique(xs)
	ys = unique(ys)

	cx = make(map[int]int, len(xs))
	cy = make(map[int]int, len(ys))

	for i, v := range xs {
		cx[v] = i
	}
	for i, v := range ys {
		cy[v] = i
	}

	return cx, cy, xs, ys
}

func markOutside(grid [][]byte) {
	R, C := len(grid), len(grid[0])

	visited := make([][]bool, R)
	for r := range visited {
		visited[r] = make([]bool, C)
	}

	type cell struct{ r, c int }
	queue := make([]cell, 0, R*C)

	push := func(r, c int) {
		if r < 0 || r >= R || c < 0 || c >= C {
			return
		}
		if !visited[r][c] && grid[r][c] == '.' {
			visited[r][c] = true
			grid[r][c] = 'o'
			queue = append(queue, cell{r, c})
		}
	}

	// enqueue all boundary '.' cells
	for c := range C {
		push(0, c)
		push(R-1, c)
	}
	for r := range R {
		push(r, 0)
		push(r, C-1)
	}

	// BFS
	for i := 0; i < len(queue); i++ {
		r, c := queue[i].r, queue[i].c
		push(r-1, c)
		push(r+1, c)
		push(r, c-1)
		push(r, c+1)
	}
}

func fillInterior(grid [][]byte) {
	for r := range len(grid) {
		for c := range len(grid[0]) {
			switch grid[r][c] {
			case '.':
				grid[r][c] = '#'
			case 'o':
				grid[r][c] = '.'
			}
		}
	}
}

func part2(input []string) int {
	coords := getCoords(input)
	cx, cy, xs, ys := compress(coords)

	R, C := len(ys), len(xs)
	grid := make([][]byte, R)
	for r := range R {
		grid[r] = make([]byte, C)
		for c := range C {
			grid[r][c] = '.'
		}
	}

	n := len(coords)
	for i := range n {
		x1, y1 := coords[i][0], coords[i][1]
		x0, y0 := coords[(i+n-1)%n][0], coords[(i+n-1)%n][1]

		cx1, cy1 := cx[x1], cy[y1]
		cx0, cy0 := cx[x0], cy[y0]

		grid[cy1][cx1] = '#'

		if x0 == x1 {
			if cy0 > cy1 {
				cy0, cy1 = cy1, cy0
			}
			for r := cy0; r <= cy1; r++ {
				grid[r][cx0] = '#'
			}
		} else if y0 == y1 {
			if cx0 > cx1 {
				cx0, cx1 = cx1, cx0
			}
			for c := cx0; c <= cx1; c++ {
				grid[cy0][c] = '#'
			}
		}
	}

	// fill polygon interior
	markOutside(grid)
	fillInterior(grid)

	// find max rectangle
	output := 0
	for i := range n {
		for j := i + 1; j < n; j++ {
			x1, y1 := coords[i][0], coords[i][1]
			x2, y2 := coords[j][0], coords[j][1]
			cx1, cy1 := cx[x1], cy[y1]
			cx2, cy2 := cx[x2], cy[y2]

			minX, maxX := cx1, cx2
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			minY, maxY := cy1, cy2
			if minY > maxY {
				minY, maxY = maxY, minY
			}

			valid := true
			for r := minY; r <= maxY && valid; r++ {
				for c := minX; c <= maxX; c++ {
					if grid[r][c] != '#' {
						valid = false
						break
					}
				}
			}
			if !valid {
				continue
			}

			area := rectArea(xs[minX], ys[minY], xs[maxX], ys[maxY])
			if area > output {
				output = area
			}
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
