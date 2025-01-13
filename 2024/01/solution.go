// --- Day 1: Historian Hysteria ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	var left []int
	var right []int
	for scanner.Scan() {
		text := scanner.Text()
		input := strings.Split(strings.TrimSpace(text), "   ")
		numStr := input[0]

		// left side conversion
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		left = append(left, num)

		// right side conversion
		numStr = input[1]
		num, err = strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		right = append(right, num)
	}

	// sort both in ascending order
	slices.Sort(left)
	slices.Sort(right)

	// calculating the distance
	// total += num1 - num2
	var total int
	for i, num1 := range left {
		var distance int
		num2 := right[i]

		if num1 > num2 {
			distance = num1 - num2
		} else {
			distance = num2 - num1
		}

		total += distance
	}

	fmt.Println("Part 1:", total)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	part1(file)
}
