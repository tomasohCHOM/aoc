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

func bfs1(r, c int, char byte, seen map[coords]bool, input []string) (int, int) {
	M, N := len(input), len(input[0])
	q := list.New()
	q.PushBack([2]int{r, c})
	area, perimeter := 0, 0

	for q.Len() > 0 {
		curr := q.Remove(q.Front()).([2]int)
		r, c := curr[0], curr[1]
		if r < 0 || c < 0 || r >= M || c >= N || input[r][c] != char {
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

func bfs2(r, c int, char byte, seen map[coords]bool, input []string) (int, int) {
	M, N := len(input), len(input[0])
	q := list.New()
	q.PushBack([2]int{r, c})
	area, edges := 0, 0

	for q.Len() > 0 {
		curr := q.Remove(q.Front()).([2]int)
		r, c := curr[0], curr[1]
		if r < 0 || c < 0 || r >= M || c >= N || input[r][c] != char {
			edges++
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
	return area, edges
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
