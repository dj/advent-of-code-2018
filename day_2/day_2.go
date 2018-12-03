package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"strings"
)

func main() {
	content, err :=	 ioutil.ReadFile("day_2.txt")
	if err != nil {
		panic(err)
	}
	input := string(content[:])

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
	fmt.Println("checksum: ", has2 * has3)
}
