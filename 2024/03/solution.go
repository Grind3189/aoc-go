// --- Day 3: Mull It Over ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	var total int
	for scanner.Scan() {
		text := scanner.Text()
		// match a string with "mul(nums,nums)"
		re := regexp.MustCompile(`mul[(]\d+,\d+[)]`)

		for _, match := range re.FindAllString(text, -1) {
			// find the num in "mul(nums,nums)"
			reg := regexp.MustCompile(`\d+`)
			strs := reg.FindAllString(match, -1)

			num1, err := strconv.Atoi(strs[0])
			if err != nil {
				panic(err)
			}
			num2, err := strconv.Atoi(strs[1])
			if err != nil {
				panic(err)
			}

			total += num1 * num2
		}
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
