package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	cost, r, c int
}

type MinHeap []Cell

func (minHeap MinHeap) Len() int            { return len(minHeap) }
func (minHeap MinHeap) Less(i, j int) bool  { return minHeap[i].cost < minHeap[j].cost }
func (minHeap MinHeap) Swap(i, j int)       { minHeap[i], minHeap[j] = minHeap[j], minHeap[i] }
func (minHeap *MinHeap) Push(x interface{}) { *minHeap = append(*minHeap, x.(Cell)) }
func (minHeap *MinHeap) Pop() interface{} {
	old := *minHeap
	N := len(old)
	cell := old[N-1]
	*minHeap = old[:N-1]
	return cell
}

func dijkstras(grid [][]int) int {
	M, N := len(grid), len(grid[0])

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, Cell{0, 0, 0})
	seen := map[[2]int]bool{}

	for minHeap.Len() > 0 {
		curr := heap.Pop(minHeap).(Cell)
		cost, r, c := curr.cost, curr.r, curr.c

		if r == M-1 && c == N-1 {
			return cost
		}
		if seen[[2]int{r, c}] {
			continue
		}
		seen[[2]int{r, c}] = true
		for _, newCoords := range [][]int{{r - 1, c}, {r + 1, c}, {r, c - 1}, {r, c + 1}} {
			nr, nc := newCoords[0], newCoords[1]
			if nr < 0 || nc < 0 || nr >= M || nc >= N || grid[nr][nc] == 1 {
				continue
			}
			newCost := cost + 1
			heap.Push(minHeap, Cell{newCost, nr, nc})
		}
	}
	return -1
}

func part1(input []string) int {
	grid := [][]int{}
	for r := 0; r < 71; r++ {
		grid = append(grid, []int{})
		for c := 0; c < 71; c++ {
			grid[r] = append(grid[r], 0)
		}
	}
	for i, line := range input {
		if i == 1024 {
			break
		}
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		grid[y][x] = 1
	}
	return dijkstras(grid)
}

func part2(input []string) (int, int) {
	grid := [][]int{}
	for r := 0; r < 71; r++ {
		grid = append(grid, []int{})
		for c := 0; c < 71; c++ {
			grid[r] = append(grid[r], 0)
		}
	}
	for i, line := range input {
		if i == 1024 {
			break
		}
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		grid[y][x] = 1
	}
	resX, resY := -1, -1
	for i := 1024; i < len(input); i++ {
		coords := strings.Split(input[i], ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		grid[y][x] = 1
		if dijkstras(grid) == -1 {
			resX, resY = x, y
			break
		}
	}
	return resX, resY
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
	r, c := part2(input)
	fmt.Printf("Part 2: (%d,%d)\n", r, c)
}
