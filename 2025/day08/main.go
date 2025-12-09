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
	Dist float64
	I    int
	J    int
	P1   []int
	P2   []int
}

func euclidean(p1, p2 []int) float64 {
	sumSqDiff := 0.0
	for i := range p1 {
		diff := float64(p1[i] - p2[i])
		sumSqDiff += diff * diff
	}
	return math.Sqrt(sumSqDiff)
}

func part1(input []string) int {
	distances := []Distance{}
	for i := range input {
		for j := i + 1; j < len(input); j++ {
			ps1 := strings.Split(input[i], ",")
			ps2 := strings.Split(input[j], ",")
			p1, p2 := []int{}, []int{}
			for co := range ps1 {
				co1, _ := strconv.Atoi(ps1[co])
				co2, _ := strconv.Atoi(ps2[co])
				p1 = append(p1, co1)
				p2 = append(p2, co2)
			}
			distances = append(distances, Distance{
				Dist: euclidean(p1, p2),
				I:    i,
				J:    j,
			})
		}
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Dist < distances[j].Dist
	})
	uf := NewUnionFind(len(input))

	for i := range input {
		dist := distances[i]
		uf.Union(dist.I, dist.J)
	}

	sizes := []int{}
	for i := range input {
		if uf.Find(i) == i {
			sizes = append(sizes, uf.GetSetSize(i))
		}
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})
	return sizes[0] * sizes[1] * sizes[2]
}

func part2(input []string) int {
	distances := []Distance{}
	for i := range input {
		for j := i + 1; j < len(input); j++ {
			ps1 := strings.Split(input[i], ",")
			ps2 := strings.Split(input[j], ",")
			p1, p2 := []int{}, []int{}
			for co := range ps1 {
				co1, _ := strconv.Atoi(ps1[co])
				co2, _ := strconv.Atoi(ps2[co])
				p1 = append(p1, co1)
				p2 = append(p2, co2)
			}
			distances = append(distances, Distance{
				Dist: euclidean(p1, p2),
				I:    i,
				J:    j,
				P1:   p1,
				P2:   p2,
			})
		}
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Dist < distances[j].Dist
	})
	uf := NewUnionFind(len(input))
	last := []int{}
	for i := range distances {
		dist := distances[i]
		if uf.Union(dist.I, dist.J) {
			last = []int{dist.P1[0], dist.P2[0]}
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
