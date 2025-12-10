package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Distance struct {
	Dist   float64
	I, J   int
	P1, P2 []int
}

func parsePoint(s string) []int {
	parts := strings.Split(s, ",")
	p := make([]int, len(parts))
	for i := range parts {
		v, _ := strconv.Atoi(parts[i])
		p[i] = v
	}
	return p
}

func euclidean(a, b []int) float64 {
	sum := 0.0
	for i := range a {
		d := float64(a[i] - b[i])
		sum += d * d
	}
	return math.Sqrt(sum)
}

func getDistances(points [][]int, includePoints bool) []Distance {
	d := []Distance{}
	N := len(points)
	for i := range N {
		for j := i + 1; j < N; j++ {
			entry := Distance{
				Dist: euclidean(points[i], points[j]),
				I:    i,
				J:    j,
			}
			if includePoints {
				entry.P1 = points[i]
				entry.P2 = points[j]
			}
			d = append(d, entry)
		}
	}
	sort.Slice(d, func(a, b int) bool {
		return d[a].Dist < d[b].Dist
	})
	return d
}

func part1(input []string) int {
	points := make([][]int, len(input))
	for i, line := range input {
		points[i] = parsePoint(line)
	}

	dists := getDistances(points, false)
	uf := NewUnionFind(len(points))
	for i := range points {
		d := dists[i]
		uf.Union(d.I, d.J)
	}

	var sizes []int
	for i := range points {
		if uf.Find(i) == i {
			sizes = append(sizes, uf.GetSetSize(i))
		}
	}

	sort.Slice(sizes, func(a, b int) bool { return sizes[a] > sizes[b] })
	return sizes[0] * sizes[1] * sizes[2]
}

func part2(input []string) int {
	points := make([][]int, len(input))
	for i, line := range input {
		points[i] = parsePoint(line)
	}
	dists := getDistances(points, true)
	uf := NewUnionFind(len(points))

	last := []int{}
	for _, d := range dists {
		if uf.Union(d.I, d.J) {
			last = []int{d.P1[0], d.P2[0]}
		}
	}
	return last[0] * last[1]
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
