// --- Day 3: Perfectly Spherical Houses in a Vacuum ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	x, y, visited := 0, 0, 1
	loc := "0-0"
	locMap := map[string]int{loc: 1}

	for scanner.Scan() {
		text := scanner.Text()

		fmt.Println(text)

		for _, runes := range text {
			char := string(runes)

			switch char {
			case "^":
				x++
			case "v":
				x--
			case ">":
				y++
			case "<":
				y--
			default:
				fmt.Println("hello")
			}

			newLoc := strconv.Itoa(x) + "-" + strconv.Itoa(y)
			locVal := locMap[newLoc]
			if locVal == 0 {
				visited++
			}
			locMap[newLoc] = locVal + 1
			loc = newLoc
		}

		fmt.Println("---")
	}

	fmt.Println("Part 1:", visited)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	part1(file)
}
