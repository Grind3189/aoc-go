// --- Day 5: Doesn't He Have Intern-Elves For This? ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	var total int
scan:
	for scanner.Scan() {
		text := scanner.Text()

		// at least 3 vowels
		{
			re := regexp.MustCompile("[aeiou]")
			matched := re.FindAllString(text, -1)
			if len(matched) < 3 {
				continue scan
			}
		}
		// one letter appeared twice in a row and
		// does not contain the blacklisted strings
		{
			blacklisted := []string{"ab", "cd", "pq", "xy"}
			hasPassed := false
			for i := 0; i < len(text)-1; i++ {
				c1 := string(text[i])
				c2 := string(text[i+1])
				str := c1 + c2

				for _, blacklist := range blacklisted {
					if blacklist == str {
						continue scan
					}
				}

				if c1 == c2 && hasPassed == false {
					hasPassed = true
				}
			}

			if hasPassed == false {
				continue scan
			}
		}

		total++

	}

	fmt.Println("Part 1:", total)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	part1(file)
}
