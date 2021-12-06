package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func NewPoint(input string) (Point, error) {
	parts := strings.Split(input, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return Point{}, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return Point{}, err
	}
	return Point{x, y}, nil
}

type Line struct {
	Start  Point
	End    Point
	Points []Point
}

func NewLine(input string) (Line, error) {
	parts := strings.Split(input, " -> ")
	start, err := NewPoint(parts[0])
	if err != nil {
		return Line{}, err
	}

	end, err := NewPoint(parts[1])
	if err != nil {
		return Line{}, err
	}

	return Line{
		Start:  start,
		End:    end,
		Points: make([]Point, 0),
	}, nil
}

func (l Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.Start.X, l.Start.Y, l.End.X, l.End.Y)
}

type Floor struct {
	Counts [100][100]int
}

func (f Floor) String() string {
	s := ""
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if f.Counts[i][j] == 0 {
				s += "."
			} else {
				s += fmt.Sprintf("%d")
			}
		}
		s += "\n"
	}
	return s
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

func main() {
	inputLines := getLines("input.txt")
	lines := make([]Line, len(inputLines))
	floor := Floor{}
	for i, line := range inputLines {
		l, err := NewLine(line)
		if err != nil {
			log.Fatalf("couldn't parse line: %v\n", err)
		}
		lines[i] = l
	}
	for _, l := range lines {
		fmt.Println(l)
	}
	log.Println(len(lines))
	fmt.Println(floor)
}
