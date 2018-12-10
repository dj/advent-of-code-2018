package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

func part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	has2 := 0
	has3 := 0
	for scanner.Scan() {
		line := scanner.Text()
		seen := make(map[rune]int)
		for _, rune := range line {
			seen[rune] = seen[rune] + 1
		}

		runesChecked := make(map[rune]bool)
		incrementedHas2 := false
		incrementedHas3 := false
		for rune, count := range seen {
			checked := runesChecked[rune]
			if !checked {
				switch count {
				case 2:
					if !incrementedHas2 {
						has2++
						incrementedHas2 = true
					}
				case 3:
					if !incrementedHas3 {
						has3++
						incrementedHas3 = true
					}
				}
			}
			runesChecked[rune] = true
		}
	}
	return has2 * has3
}

func part2(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	lineCount := 0
	lines := make(map[int]string)

	for scanner.Scan() {
		line := scanner.Text()
		lines[lineCount] = line
		lineCount = lineCount + 1
	}

	for _, line1 := range lines {
		for _, line2 := range lines {
			mismatches := 0

			for i := 0; i < len(line1); i++ {
				if line1[i] != line2[i] {
					mismatches++
				}
				if mismatches >= 2 {
					break
				}
			}
			if mismatches == 1 {
				common := ""
				for i := 0; i < len(line1); i++ {
					if line1[i] == line2[i] {
						common = common + string(line1[i])
					}
				}
				return common
			}
		}
	}

	return ""
}

func main() {
	content, err := ioutil.ReadFile("day_2.txt")
	if err != nil {
		panic(err)
	}
	input := string(content[:])

	fmt.Println("part 1: ", part1(input))
	fmt.Println("part 2: ", part2(input))
}
