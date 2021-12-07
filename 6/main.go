package main

import (
	"bufio"
	"fmt"
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

func Day(numbers []int) []int {
	newFish := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == 0 {
			numbers[i] = 6
			newFish++
			continue
		}

		numbers[i]--
	}
	for i := 0; i < newFish; i++ {
		numbers = append(numbers, 8)
	}
	return numbers
}

func DayMap(fish map[int]int) map[int]int {
	newMap := make(map[int]int)
	for k, v := range fish {
		switch k {
		case 0:
			newMap[6] += v
			newMap[8] = v
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 5:
			fallthrough
		case 6:
			fallthrough
		case 7:
			fallthrough
		case 8:
			newMap[k-1] += v
		default:
			log.Fatalf("invalid count value: %v\n", k)
		}
	}
	return newMap
}

// 2140412022350 guess for part 2
// 1631647919273 actual answer
func main() {
	inputLines := getLines("input.txt")
	numbers := parseNumbers(inputLines[0])
	log.Println(len(numbers))
	fish := make(map[int]int)
	for _, v := range numbers {
		fish[v]++
	}

	for i := 0; i < 256; i++ {
		fish = DayMap(fish)
	}

	fmt.Println(fish)

	sum := 0
	for _, v := range fish {
		sum += v
	}

	fmt.Println(sum)

}
