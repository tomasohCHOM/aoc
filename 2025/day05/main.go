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

func parseIdsAndRanges(input []string) ([]int, [][]int) {
	idRanges, i := [][]int{}, 0
	for i < len(input) {
		if input[i] == "" {
			i++
			break
		}
		r := strings.Split(input[i], "-")
		lower, _ := strconv.Atoi(r[0])
		upper, _ := strconv.Atoi(r[1])
		idRanges = append(idRanges, []int{lower, upper})
		i++
	}
	ids := []int{}
	for i < len(input) {
		id, _ := strconv.Atoi(input[i])
		ids = append(ids, id)
		i++
	}
	return ids, idRanges
}

func part1(ids []int, idRanges [][]int) int {
	output := 0
	for _, id := range ids {
		for _, r := range idRanges {
			if id >= r[0] && id <= r[1] {
				output++
				break
			}
		}
	}
	return output
}

func part2(idRanges [][]int) int {
	sort.Slice(idRanges, func(i, j int) bool {
		if idRanges[i][0] != idRanges[j][0] {
			return idRanges[i][0] < idRanges[j][0]
		}
		return idRanges[i][1] < idRanges[j][1]
	})

	merged := [][]int{}
	for _, r := range idRanges {
		if len(merged) == 0 {
			merged = append(merged, r)
			continue
		}
		last := merged[len(merged)-1]
		if r[0] <= last[1]+1 {
			if r[1] > last[1] {
				last[1] = r[1]
			}
		} else {
			merged = append(merged, r)
		}
	}
	output := 0
	for _, r := range merged {
		output += r[1] - r[0] + 1
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
	ids, idRanges := parseIdsAndRanges(input)
	fmt.Println("Part 1:", part1(ids, idRanges))
	fmt.Println("Part 2:", part2(idRanges))
}
