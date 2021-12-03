package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type DiagReport struct {
	Lines       []string
	lcbFiltered []string
	mcbFiltered []string
}

func (d *DiagReport) Load(filename string) {
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
	d.Lines = lines
}

func getMCB(lines []string, index int) rune {
	one := 0
	zero := 0
	for _, line := range lines {
		switch line[index] {
		case '1':
			one++
		case '0':
			zero++
		}
	}
	if one >= zero {
		return '1'
	} else {
		return '0'
	}
}

func (d DiagReport) getLCB(lines []string, index int) rune {
	one := 0
	zero := 0
	for _, line := range lines {
		switch line[index] {
		case '1':
			one++
		case '0':
			zero++
		}
	}
	if zero <= one {
		return '0'
	} else {
		return '1'
	}
}

func (d *DiagReport) getOxygen() int64 {
	d.mcbFiltered = d.Lines
	for i := 0; i < 12; i++ {
		mcb := getMCB(d.mcbFiltered, i)
		d.mcbFiltered = filterBit(d.mcbFiltered, i, mcb)
		if len(d.mcbFiltered) == 1 {
			break
		}
	}
	oxygen, err := strconv.ParseInt(d.mcbFiltered[0], 2, 64)
	if err != nil {
		log.Fatalf("can't parse oxygen: %v\n", err)
	}
	return oxygen
}

func (d *DiagReport) getCarbon() int64 {
	d.lcbFiltered = d.Lines
	for i := 0; i < 12; i++ {
		lcb := d.getLCB(d.lcbFiltered, i)
		d.lcbFiltered = filterBit(d.lcbFiltered, i, lcb)

		if len(d.lcbFiltered) == 1 {
			break
		}

	}
	carbon, err := strconv.ParseInt(d.lcbFiltered[0], 2, 64)
	if err != nil {
		log.Fatalf("can't parse carbon: %v\n", err)
	}
	return carbon
}

func filterBit(lines []string, index int, bit rune) []string {
	filtered := make([]string, 0)
	for _, line := range lines {
		if rune(line[index]) == bit {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

func main() {
	diag := DiagReport{}
	diag.Load("input.txt")
	oxygen := diag.getOxygen()
	carbon := diag.getCarbon()

	log.Println(oxygen)
	log.Println(carbon)
	log.Println(oxygen * carbon)
}

// func main() {
// 	file, err := os.Open("./input.txt")
// 	if err != nil {
// 		log.Fatalf("read file: %v\n", err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)

// 	counter := make([]BitCounter, 12)
// 	for scanner.Scan() {
// 		for i, v := range scanner.Text() {
// 			switch v {
// 			case '0':
// 				counter[i].Zero++
// 			case '1':
// 				counter[i].One++
// 			}
// 		}
// 	}
// 	mcb := ""
// 	lcb := ""
// 	for _, counter := range counter {
// 		mcb += fmt.Sprintf("%d", counter.getMCB())
// 		lcb += fmt.Sprintf("%d", counter.getLCB())
// 	}
// 	log.Println(mcb)
// 	log.Println(lcb)
// }
