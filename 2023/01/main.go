package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		panic("Unable to open file")
	}
	defer data.Close()

	var total int
	var numToAdd []int
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		var numStrs []string
		text := scanner.Text()

		// get all available numbers in a line of text
		// put it in an array
		for _, char := range text {
			// skip if the unicode is not a number
			if unicode.IsNumber(char) == false {
				continue
			}

			// convert the unicode into string and append it
			str := string(char)
			numStrs = append(numStrs, str)
		}

		// get first and last string of number in the array and concatinate it
		// e.g [1,3]
		// concat = 13
		num1 := numStrs[0]
		num2 := numStrs[len(numStrs)-1]
		concat := num1 + num2

		// convert the number into string
		// append it to the list of numbers to add
		num, err := strconv.Atoi(concat)
		if err != nil {
			panic("Unable to convert string to number")
		}
		numToAdd = append(numToAdd, num)
	}

	// get the sum of all numbers
	for _, num := range numToAdd {
		total += num
	}
}
