package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Values        [5][5]int
	Marked        [5][5]bool
	ColumnCount   [5]int
	RowCount      [5]int
	WinningNumber int
}

func (b Board) String() string {
	s := ""
	for _, row := range b.Values {
		s += fmt.Sprintf("%v\n", row)
	}
	for _, row := range b.Marked {
		s += fmt.Sprintf("%v\n", row)
	}
	s += fmt.Sprintln(b.ColumnCount)
	s += fmt.Sprintln(b.RowCount)
	return s
}

func NewBoard(values [5][5]int) Board {
	return Board{
		Values:      values,
		Marked:      [5][5]bool{},
		ColumnCount: [5]int{},
		RowCount:    [5]int{},
	}
}

func (b *Board) Play(number int) bool {
	for i, row := range b.Values {
		for j, v := range row {
			if number == v {
				b.Marked[i][j] = true
				b.ColumnCount[j]++
				b.RowCount[i]++

				if b.ColumnCount[j] == 5 || b.RowCount[i] == 5 {
					b.WinningNumber = number
					return true
				}
				return false
			}
		}
	}
	return false
}

func (b Board) Sum(marked bool) int {
	sum := 0
	for i, row := range b.Values {
		for j, v := range row {
			if b.Marked[i][j] == marked {
				sum += v
			}
		}
	}
	return sum
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
	nums := make([]int, 0)
	parts := strings.Split(input, ",")
	for _, v := range parts {
		val, _ := strconv.Atoi(v)
		nums = append(nums, val)
	}
	return nums
}

func parseBoard(rows [5]string) Board {
	values := [5][5]int{}
	for i, row := range rows {
		row = strings.Trim(row, " ")
		row = strings.ReplaceAll(row, "  ", " ")
		temp := parseBoardRow(row)
		values[i] = temp
	}
	b := NewBoard(values)
	return b
}

func parseBoardRow(row string) [5]int {
	parts := strings.Split(row, " ")
	values := [5]int{}
	for i, v := range parts {
		val, _ := strconv.Atoi(v)
		values[i] = val
	}
	return values
}

func main() {
	lines := getLines("input.txt")
	numbers := parseNumbers(lines[0])
	boards := make([]Board, 0)
	for i := 2; i < len(lines); i += 6 {
		temp := [5]string{}
		copy(temp[:], lines[i:i+5])
		b := parseBoard(temp)
		boards = append(boards, b)
	}

	for _, v := range numbers {
		for i := range boards {
			if boards[i].Play(v) {
				log.Println(boards[i].Sum(false) * boards[i].WinningNumber)
				return
			}
		}
	}
	// for _, b := range boards {
	// 	fmt.Println(b)
	// }
}
