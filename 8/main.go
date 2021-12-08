package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Note struct {
	Segments []string
	Output   []string
}

func ParseNote(input string) Note {
	parts := strings.Split(input, " | ")
	return Note{
		Segments: strings.Split(parts[0], " "),
		Output:   strings.Split(parts[1], " "),
	}
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

func main() {
	inputLines := getLines("input.txt")
	notes := make([]Note, 0)
	count := 0
	for _, line := range inputLines {
		note := ParseNote(line)
		notes = append(notes, note)
		for _, v := range note.Output {
			length := len(v)
			if length == 2 || length == 4 || length == 3 || length == 7 {
				count++
			}
		}
	}
	fmt.Println(notes)
	log.Println(count)
}
