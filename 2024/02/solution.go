// --- Day 2: Red-Nosed Reports ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	var safe int
outer:
	for scanner.Scan() {
		const MIN, MAX = 1, 3

		text := scanner.Text()
		input := strings.Split(text, " ")

		asc, desc, prev := false, false, 0
	check:
		for i, txt := range input {
			diff := 0
			num, err := strconv.Atoi(txt)
			if err != nil {
				panic(err)
			}

			if i == 0 {
				prev = num
				continue check
			}

			// each outer loop should only be asc or desc
			// to be considered valid
			if num > prev && desc == false {
				asc = true
				diff = num - prev
			} else if num < prev && asc == false {
				desc = true
				diff = prev - num
			} else {
				continue outer
			}

			prev = num

			// the difference between each level should be between
			// the minimum and maximum value to be considered valid
			if diff < MIN || diff > MAX {
				continue outer
			}

		}
		safe++
	}

	fmt.Println("Part 1:", safe)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	part1(file)
}
