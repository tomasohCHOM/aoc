package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getAdjacencyList(lines []string) map[int]map[int]bool {
	adj := make(map[int]map[int]bool)
	for _, line := range lines {
		mapping := strings.Split(line, "|")
		u, _ := strconv.Atoi(mapping[0])
		v, _ := strconv.Atoi(mapping[1])
		if adj[u] == nil {
			adj[u] = make(map[int]bool)
		}
		adj[u][v] = true
	}
	return adj
}

func getUpdateLists(lines []string) [][]int {
	updateLists := [][]int{}
	for _, line := range lines {
		list := strings.Split(line, ",")
		updateList := []int{}
		for _, v := range list {
			elem, _ := strconv.Atoi(v)
			updateList = append(updateList, elem)
		}
		updateLists = append(updateLists, updateList)
	}
	return updateLists
}

func parseInput(input []string) (map[int]map[int]bool, [][]int) {
	var middle int
	for middle = 0; middle < len(input); middle++ {
		if input[middle] == "" {
			break
		}
	}
	adj := getAdjacencyList(input[:middle])
	updateLists := getUpdateLists(input[middle+1:])
	return adj, updateLists
}

func isCorrect(adj map[int]map[int]bool, updateList []int) bool {
	for i := 0; i < len(updateList); i++ {
		for j := i + 1; j < len(updateList); j++ {
			if adj[updateList[j]][updateList[i]] {
				return false
			}
		}
	}
	return true
}

func part1(input []string) int {
	adj, updateLists := parseInput(input)
	output := 0
	for _, updateList := range updateLists {
		if isCorrect(adj, updateList) {
			output += updateList[len(updateList)/2]
		}
	}
	return output
}

func part2(input []string) int {
	adj, updateLists := parseInput(input)
	incorrects := [][]int{}
	for _, updateList := range updateLists {
		if !isCorrect(adj, updateList) {
			incorrects = append(incorrects, updateList)
		}
	}
	output := 0
	for _, incorrect := range incorrects {
		sort.Slice(incorrect, func(i, j int) bool {
			a := incorrect[i]
			b := incorrect[j]
			return adj[a][b]
		})
		output += incorrect[len(incorrect)/2]
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
