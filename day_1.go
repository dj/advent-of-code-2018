package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	frequency := 0
	for scanner.Scan() {
		line := scanner.Text()
		op := line[0]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		switch op {
		case '+':
			frequency = frequency + num
		case '-':
			frequency = frequency - num
		}
	}
	return frequency
}

func part2(input string) int {
	frequency := 0
	seen := make(map[int]bool)
	seen[frequency] = true
	var seenTwice *int
	// loop through the input until we see the same freq twice
	for {
		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {
			if seenTwice != nil {
				return *seenTwice
			}
			line := scanner.Text()
			op := line[0]
			num, err := strconv.Atoi(line[1:])
			if err != nil {
				panic(err)
			}
			switch op {
			case '+':
				frequency = frequency + num
			case '-':
				frequency = frequency - num
			}
			if seenTwice == nil && seen[frequency] == true {
				seenTwice = &frequency
			}
			seen[frequency] = true
		}
	}
}

func main() {
	input := readInput("input/day_1.txt")
	fmt.Println("part 1: ", part1(input))
	fmt.Println("part 2: ", part2(input))
}