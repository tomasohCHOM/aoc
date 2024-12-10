package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getDisk(input string) []string {
	disk, id := []string{}, 0
	for i, c := range input {
		val, _ := strconv.Atoi(string(c))
		if i%2 == 1 {
			for i := 0; i < val; i++ {
				disk = append(disk, ".")
			}
		} else {
			for i := 0; i < val; i++ {
				disk = append(disk, strconv.Itoa(id))
			}
			id++
		}
	}
	return disk
}

func checkSum(disk []string) int {
	sum := 0
	for i, c := range disk {
		if disk[i] == "." {
			continue
		}
		val, _ := strconv.Atoi(c)
		sum += i * val
	}
	return sum
}

func part1(input string) int {
	disk := getDisk(input)
	l, r := 0, len(disk)-1
	for l < r {
		for l < r && disk[l] != "." {
			l++
		}
		for l < r && disk[r] == "." {
			r--
		}
		if l < r {
			disk[l], disk[r] = disk[r], disk[l]
		}
	}
	return checkSum(disk)
}

func part2(input string) int {
	disk := getDisk(input)
	r := len(disk) - 1
	for r >= 0 {
		for r >= 0 && disk[r] == "." {
			r--
		}
		if r < 0 {
			break
		}
		char, dr := disk[r], r
		for dr > 0 && disk[dr-1] == char {
			dr--
		}
		sizeR := r - dr + 1

		l, found := 0, false
		for l < r {
			for l < r && disk[l] != "." {
				l++
			}
			if l+sizeR-1 >= r {
				break
			}
			canFit := true
			for i := 0; i < sizeR; i++ {
				if disk[l+i] != "." {
					canFit = false
					break
				}
			}
			if canFit {
				found = true
				break
			}
			l++
		}
		if found {
			for i := l; i < l+sizeR; i++ {
				disk[i] = char
				disk[r] = "."
				r--
			}
		}
		r = dr - 1
	}
	return checkSum(disk)
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
	fmt.Println("Part 1:", part1(input[0]))
	fmt.Println("Part 2:", part2(input[0]))
}
