package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type SlidingWindowInt struct {
	Size   int
	Values []int
}

func NewSlidingWindowInt(size int) SlidingWindowInt {
	return SlidingWindowInt{
		Size:   size,
		Values: make([]int, 0),
	}
}

func (swi SlidingWindowInt) String() string {
	// implement Stringer
	return fmt.Sprintf("%v: %d", swi.Values, swi.Sum())
}

func (swi SlidingWindowInt) Sum() int {
	// range over values and return sum
	sum := 0
	for _, v := range swi.Values {
		sum += v
	}
	return sum
}

func (swi SlidingWindowInt) Full() bool {
	// if the len of the values is the size, then it is full
	return len(swi.Values) == swi.Size
}

func (swi *SlidingWindowInt) Append(value int) {
	// append if buffer isn't full
	if len(swi.Values) < swi.Size {
		swi.Values = append(swi.Values, value)
	} else {
		// if full, overwrite the first value and append new value
		copy(swi.Values, swi.Values[1:])
		swi.Values[swi.Size-1] = value
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("read file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cur, prev int
	count := 0
	sw := NewSlidingWindowInt(3)
	for scanner.Scan() {
		cur, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("strconv: %v\n", err)
		}
		sw.Append(cur)
		fmt.Println(sw)

		// if sliding window is full
		if sw.Full() {
			// if there wasn't a previous measurement
			if prev == 0 {
				// set prev to the current measurement
				prev = sw.Sum()
				continue
			}

			// save current measurement
			cur = sw.Sum()

			// if current is greater than previous
			if cur > prev {
				// increment
				count++
			}

			// set previous to current
			prev = cur
		}

	}
	log.Println(count)
}
