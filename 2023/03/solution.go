// --- Day 3: Gear Ratios ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)

	// create an array of string
	// modify the each text, add '.' at beginning and end
	// this wont affect the value of each since '.' is not a symbol according to instruction
	// now we'll always have a value at the left and right side of the index
	// "123..%123" to ".123..%123."
	var input []string
	for scanner.Scan() {
		text := "." + scanner.Text() + "."
		input = append(input, text)
	}

	var total int
	for i, text := range input {

		// we'll use numStr for appending numbers
		// e.g. "12" + "3" "123"
		// then we will convert this later to actual number if this is a part number
		// we'll use txtToCheck to store every character sorrounding the number
		// then we will check it at once
		var numStr, txtToCheck string
		var start int
	inner:
		for j, rune := range text {

			if unicode.IsDigit(rune) {
				// store the starting index of the 1st number found
				if numStr == "" {
					start = j - 1
				}

				numStr += string(rune)
				continue inner
			}

			// at this point the current character in the loop is not a number
			// check if we have a value of number strings before
			// performing checks
			if numStr != "" {
				var before, after, up, down string
				valid := false

				// get left and right value of numbers
				// e.g. ".123*" = ".*"
				before = string(text[start])
				after = string(text[j])

				// get text row before this current text
				// e.g ".*..$" to ".*..$" // get this
				//     ".123."            // current text
				if i > 0 {
					up = input[i-1][start : j+1]
				}

				// get the bottom text after this text
				// e.g. ".123." // current text
				//      "....." // get this
				if i < len(input)-1 {
					down = input[i+1][start : j+1]
				}

				// append it all before checking
				txtToCheck = before + after + up + down
			check:
				for _, char := range txtToCheck {
					if char != '.' {
						valid = true
						break check
					}
				}

				if valid {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						panic("Unable to convert string to number")
					}
					total += num
				}

				// reset every operation
				numStr, txtToCheck, before, up, down, after = "", "", "", "", "", ""
				start = 0

			}
		}
	}

	fmt.Println("Total:", total)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Unable to open file")
	}

	part1(file)
}
