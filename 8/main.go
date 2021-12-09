package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	values map[string]bool
}

func SetFromString(input string) Set {
	return SetFromValues(strings.Split(input, ""))
}

func SetFromValues(values []string) Set {
	set := Set{values: make(map[string]bool)}
	for _, v := range values {
		set.values[v] = true
	}
	return set
}

func (s Set) Equals(s3 Set) bool {
	return s.Contains(s3) && s3.Contains(s)
}

func (s Set) Union(s2 Set) Set {
	s3 := Set{}
	for k, v := range s.values {
		// add if in second set
		if s2.values[k] {
			s3.values[k] = v
		}
	}
	return s3
}

func (s Set) Add(s2 Set) Set {
	s3 := Set{values: make(map[string]bool)}
	for k, v := range s.values {
		// add all from first set
		s3.values[k] = v
	}
	for k, v := range s2.values {
		// add all from second set
		s3.values[k] = v
	}
	return s3
}

func (s Set) Subtract(s2 Set) Set {
	s3 := Set{values: make(map[string]bool)}
	for k, v := range s.values {
		// add all from first set
		s3.values[k] = v
	}
	for k := range s2.values {
		// remove all from second set
		s3.values[k] = false
	}
	return s3
}

func (s Set) Contains(s2 Set) bool {
	for k := range s2.values {
		if s2.values[k] {
			// if in second set
			if s.values[k] {
				// and in first, continue
				continue
			} else {
				// else missing
				return false
			}
		}
	}
	return true
}

type Note struct {
	Segments       []string
	Output         []string
	TranslationMap map[string]int
	Value          string
}

func ParseNote(input string) Note {
	parts := strings.Split(input, " | ")
	note := Note{
		Segments:       strings.Split(parts[0], " "),
		Output:         strings.Split(parts[1], " "),
		TranslationMap: make(map[string]int),
		Value:          "",
	}
	var four, one Set
	for _, v := range note.Segments {
		switch len(v) {
		case 2:
			note.TranslationMap[v] = 1
			one = SetFromString(v)
		case 4:
			note.TranslationMap[v] = 4
			four = SetFromString(v)
		case 3:
			note.TranslationMap[v] = 7
		case 7:
			note.TranslationMap[v] = 8
		}
	}
	fourMinusOne := four.Subtract(one)
	for _, v := range note.Segments {
		switch len(v) {
		case 5:
			s := SetFromString(v)
			if s.Contains(one) {
				note.TranslationMap[v] = 3
			} else if s.Contains(fourMinusOne) {
				note.TranslationMap[v] = 5
			} else {
				note.TranslationMap[v] = 2
			}
		case 6:
			s := SetFromString(v)
			if s.Contains(four) {
				note.TranslationMap[v] = 9
			} else if s.Contains(fourMinusOne) {
				note.TranslationMap[v] = 6
			} else {
				note.TranslationMap[v] = 0
			}
		}
	}

	for _, v := range note.Output {
		for k := range note.TranslationMap {
			s1 := SetFromString(v)
			s2 := SetFromString(k)
			if s1.Equals(s2) {
				note.Value += fmt.Sprintf("%d", note.TranslationMap[k])
			}
		}
	}

	return note
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
	sum := 0
	for _, line := range inputLines {
		note := ParseNote(line)
		value, err := strconv.Atoi(note.Value)
		fmt.Printf("%#v\n", note)
		sum += value
		if err != nil {
			log.Fatalf("%v\n", err)
		}

	}
	log.Println(sum)
}
