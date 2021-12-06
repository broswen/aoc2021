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

func (p Point) Add(p2 Point) Point {
	return Point{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}

func (p Point) Equals(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}

type Line struct {
	Start Point
	End   Point
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
		Start: start,
		End:   end,
	}, nil
}

func (l Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.Start.X, l.Start.Y, l.End.X, l.End.Y)
}

type Floor struct {
	Counts [1000][1000]int
}

func (f Floor) String() string {
	s := ""
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
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
	// lines := make([]Line, len(inputLines))
	floor := Floor{}
	danger := 0
	for i, line := range inputLines {
		log.Println("parsing line", i)
		l, err := NewLine(line)
		log.Println(l)
		if err != nil {
			log.Fatalf("couldn't parse line: %v\n", err)
		}
		dx := l.End.X - l.Start.X
		dy := l.End.Y - l.Start.Y
		if dx > 0 {
			dx = 1
		} else if dx < 0 {
			dx = -1
		}
		if dy > 0 {
			dy = 1
		} else if dy < 0 {
			dy = -1
		}
		x := l.Start.X
		y := l.Start.Y
		for {
			if floor.Counts[y][x] == 1 {
				danger++
			}
			floor.Counts[y][x]++
			if x == l.End.X && y == l.End.Y {
				break
			}
			x += dx
			y += dy
		}
	}
	// for _, row := range floor.Counts {
	// 	for j := range row {
	// 		if row[j] > 1 {
	// 			danger++
	// 		}
	// 	}
	// }
	output, _ := os.Create("output.txt")
	defer output.Close()
	for _, row := range floor.Counts {
		output.WriteString(fmt.Sprintf("%v\n", row))
	}
	log.Println(danger)
}
