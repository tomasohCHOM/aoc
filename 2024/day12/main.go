package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type coords struct {
	r, c int
}

func inBounds(r, c int, input []string) bool {
	return r >= 0 && c >= 0 && r < len(input) && c < len(input[0])
}

func bfs1(r, c int, char byte, seen map[coords]bool, input []string) (int, int) {
	q := list.New()
	q.PushBack([2]int{r, c})
	area, perimeter := 0, 0

	for q.Len() > 0 {
		curr := q.Remove(q.Front()).([2]int)
		r, c := curr[0], curr[1]
		if !inBounds(r, c, input) || input[r][c] != char {
			perimeter++
			continue
		}
		if seen[coords{r, c}] {
			continue
		}
		seen[coords{r, c}] = true
		area++
		for _, newCoords := range [][]int{{r - 1, c}, {r + 1, c}, {r, c - 1}, {r, c + 1}} {
			nr, nc := newCoords[0], newCoords[1]
			q.PushBack([2]int{nr, nc})
		}
	}
	return area, perimeter
}

func part1(input []string) int {
	output := 0
	seen := map[coords]bool{}
	for r := 0; r < len(input); r++ {
		for c := 0; c < len(input[0]); c++ {
			area, perimeter := bfs1(r, c, input[r][c], seen, input)
			output += area * perimeter
		}
	}
	return output
}

func getEdges(r, c int, char byte, input []string) int {
	edges := 0

	isOutsideCorner := func(dr1, dc1, dr2, dc2 int) bool {
		return ((!inBounds(r+dr1, c+dc1, input) || input[r+dr1][c+dc1] != char) &&
			(!inBounds(r+dr2, c+dc2, input) || input[r+dr2][c+dc2] != char))
	}

	isInsideCorner := func(dr1, dc1, dr2, dc2 int) bool {
		return ((inBounds(r+dr1, c+dc1, input) && input[r+dr1][c+dc1] == char) &&
			(inBounds(r+dr2, c+dc2, input) && input[r+dr2][c+dc2] == char) &&
			(!inBounds(r+dr1+dr2, c+dc1+dc2, input) || input[r+dr1+dr2][c+dc1+dc2] != char))
	}

	cornerDeltas := [][]int{
		{1, 0, 0, 1},
		{1, 0, 0, -1},
		{-1, 0, 0, 1},
		{-1, 0, 0, -1},
	}
	for _, deltas := range cornerDeltas {
		dr1, dc1, dr2, dc2 := deltas[0], deltas[1], deltas[2], deltas[3]
		if isOutsideCorner(dr1, dc1, dr2, dc2) {
			edges++
		}
		if isInsideCorner(dr1, dc1, dr2, dc2) {
			edges++
		}
	}
	return edges
}

func bfs2(r, c int, char byte, seen map[coords]bool, input []string) (int, int) {
	q := list.New()
	q.PushBack([2]int{r, c})
	area, edges := 0, 0

	for q.Len() > 0 {
		curr := q.Remove(q.Front()).([2]int)
		r, c := curr[0], curr[1]
		if !inBounds(r, c, input) || input[r][c] != char {
			continue
		}
		if seen[coords{r, c}] {
			continue
		}
		seen[coords{r, c}] = true
		area++
		edges += getEdges(r, c, char, input)

		for _, newCoords := range [][]int{{r - 1, c}, {r + 1, c}, {r, c - 1}, {r, c + 1}} {
			nr, nc := newCoords[0], newCoords[1]
			q.PushBack([2]int{nr, nc})
		}
	}
	return area, edges
}

func part2(input []string) int {
	output := 0
	seen := map[coords]bool{}
	for r := 0; r < len(input); r++ {
		for c := 0; c < len(input[0]); c++ {
			area, edges := bfs2(r, c, input[r][c], seen, input)
			output += area * edges
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
	input, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
