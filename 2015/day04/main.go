package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"math"
	"os"
	"strings"
)

func findHashWithPrefix(input string, prefix string) int {
	for i := 0; i < math.MaxInt32; i++ {
		toHash := fmt.Sprintf("%s%d", input, i)
		hashed := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
		if strings.HasPrefix(hashed, prefix) {
			return i
		}
	}
	panic("No hash found")
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
	fmt.Println("Part 1:", findHashWithPrefix(input[0], "00000"))
	fmt.Println("Part 2:", findHashWithPrefix(input[0], "000000"))
}
