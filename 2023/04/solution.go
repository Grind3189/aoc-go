// --- Day 4: Scratchcards ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	var total int

	for scanner.Scan() {
		// sample text
		// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		text := scanner.Text()
		parts := strings.Split(text, ":")
		cards := strings.Split(parts[1], "|")
		// split the string into array of string
		// both for winning numbers and the numbers we have
		// now we have [41, 48, 83 ...]
		winningNums := strings.Split(strings.TrimSpace(cards[0]), " ")
		input := strings.Split(strings.TrimSpace(cards[1]), " ")

		var winningCount int
		for _, wNum := range winningNums {
			// some input are blanks which is invalid
			if wNum == "" {
				continue
			}

			for _, num := range input {
				if num == wNum {
					// we dont need the winning numbers
					// what we need is how many times we will loop
					winningCount += 1
				}
			}
		}

		// calculations
		// e.g. winningCount = 4
		// 1, (1 * 2 = 2), (2 * 2 = 4), (4 * 2 = 8)
		// answer is: 8
		var points int
		for i := 0; i < winningCount; i++ {
			multiplicator := 2
			if i == 0 {
				points = 1
				continue
			}
			points *= multiplicator

		}
		total += points
	}

	fmt.Println("Part 1:", total)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Unable to open file")
	}
	defer file.Close()

	part1(file)
}
