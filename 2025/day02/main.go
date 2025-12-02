package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(input []string) int {
	output := 0
	for _, idRange := range input {
		ids := strings.Split(idRange, "-")
		firstId, _ := strconv.Atoi(ids[0])
		lastId, _ := strconv.Atoi(ids[1])

		for id := firstId; id <= lastId; id++ {
			strID := strconv.Itoa(id)
			N := len(strID)
			if N%2 == 0 && strID[:N/2] == strID[N/2:] {
				output += id
			}
		}
	}
	return output
}

func part2(input []string) int {
	output := 0
	for _, idRange := range input {
		ids := strings.Split(idRange, "-")
		firstId, _ := strconv.Atoi(ids[0])
		lastId, _ := strconv.Atoi(ids[1])

		for id := firstId; id <= lastId; id++ {
			strID := strconv.Itoa(id)
			N := len(strID)

			for k := 1; k <= N/2; k++ {
				if N%k != 0 {
					continue
				}
				part := strID[:k]
				notValid := true

				for i := 0; i < N; i += k {
					if strID[i:i+k] != part {
						notValid = false
						break
					}
				}

				if notValid {
					output += id
					break
				}
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
	parsedInput := strings.Split(input[0], ",")

	fmt.Println("Part 1:", part1(parsedInput))
	fmt.Println("Part 2:", part2(parsedInput))
}
