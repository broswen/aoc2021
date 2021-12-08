package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	parts := strings.Split(input, ",")
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

func max(numbers []int) int {
	max := 0
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max
}
func minIndex(numbers []int) (int, int) {
	min := math.MaxInt32
	index := -1
	for i, v := range numbers {
		if v < min {
			min = v
			index = i
		}
	}
	return min, index
}

func triangle(number int) int {
	return int((math.Pow(float64(number), 2) + float64(number)) / float64(2))
}

func calcCost(numbers []int) []int {
	maxPosition := max(numbers)
	cost := make([]int, maxPosition)
	for i := range cost {
		for _, v := range numbers {
			diff := i - v
			if diff < 0 {
				cost[i] += triangle(diff * -1)
			} else {
				cost[i] += triangle(diff)
			}
		}
	}
	return cost
}

func main() {
	inputLines := getLines("input.txt")
	numbers := parseNumbers(inputLines[0])
	cost := calcCost(numbers)
	fmt.Println(cost)
	fmt.Println(minIndex(cost))
}
