package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Sub struct {
	X   int
	Y   int
	Aim int
}

func NewSub() Sub {
	return Sub{}
}

// print sub details and final multiplied position
func (s Sub) String() string {
	return fmt.Sprintf("%d, %d, %d: %d", s.X, s.Y, s.Aim, s.X*s.Y)
}

// return tuple of sub location and aim
func (s Sub) Position() (int, int, int) {
	return s.X, s.Y, s.Aim
}

// use command and amount to modify sub position
func (s *Sub) ParseCommand(command string, amount int) {
	switch command {
	case "forward":
		s.X += amount
		s.Y += (s.Aim * amount)
	case "up":
		s.Aim -= amount
	case "down":
		s.Aim += amount
	default:
		log.Fatalf("unknown command: %v\n", command)
	}

}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("read file: %v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sub := NewSub()

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		command := parts[0]
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("invalid amount: %v\n", parts[1])
		}
		sub.ParseCommand(command, amount)
	}
	log.Println(sub)
}
