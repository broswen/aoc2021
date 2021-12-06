package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Fish struct {
	Timer int
}

func (f *Fish) Live() []Fish {
	f.Timer--
	if f.Timer == 0 {
		// ready to spawn a new fish
		f.Timer = 6
		return []Fish{{Timer: 8}}
	}
	return []Fish{}
}

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

// 2140412022350 guess for part 2
func main() {
	inputLines := getLines("input.txt")
	numbers := parseNumbers(inputLines[0])
	log.Println(len(numbers))

	for i := 0; i < 256; i++ {
		// log.Println("day", i)
		numbers = Day(numbers)
	}

	log.Println(len(numbers))

}
