package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("read file: %v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseNumbers(input string) []int {
	parts := strings.Split(input, "")
	numbers := make([]int, 0)
	for _, v := range parts {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("couldn't parse int: %v\n", err)
		}
		numbers = append(numbers, val)
	}
	return numbers
}

func CheckLow(i, j int, numbers [][]int) bool {
	value := numbers[i][j]
	if i != 0 {
		// not top row
		if numbers[i-1][j] <= value {
			return false
		}
	}

	if j != 0 {
		// not left column
		if numbers[i][j-1] <= value {
			return false
		}
	}

	if i < len(numbers)-1 {
		// not bottom row
		if numbers[i+1][j] <= value {
			return false
		}
	}
	if j < len(numbers[i])-1 {
		// not right column
		if numbers[i][j+1] <= value {
			return false
		}
	}
	return true
}

func main() {
	inputLines := getLines("input.txt")
	numbers := make([][]int, 0)
	for _, line := range inputLines {
		numbers = append(numbers, parseNumbers(line))
	}
	sum := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			if CheckLow(i, j, numbers) {
				sum += 1 + numbers[i][j]
			}
		}
	}
	log.Println(sum)
}
