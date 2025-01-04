package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello")

	data, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, "fetch: %v\n", err)
	}
	defer data.Close()

	var left []int
	var right []int
	rightCount := make(map[int]int)

	// Get the data from file and seperate left and right numbers
	input := bufio.NewScanner(data)
	for input.Scan() {
		parts := strings.Fields(input.Text())

		if len(parts) != 2 {
			continue
		}

		// parsed the left side from string to int
		// append the int to the array
		lParsed, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv:%v\n", err)
		}
		// parsed the right side and append it to the arr
		rParsed, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv:%v\n", err)
		}

		left = append(left, lParsed)
		right = append(right, rParsed)
		rightCount[rParsed]++
	}

	// sort both array in asc order
	slices.Sort(left)
	slices.Sort(right)

	total := distance(left, right)
	fmt.Println("Total distance:", total)

	// calculate the similarity score of both array
	var simScore int
	for _, lVal := range left {
		simScore += lVal * rightCount[lVal]
	}

	fmt.Println("Similarity score:", simScore)
}

func distance(left, right []int) int {
	// calculate the distance of the two array
	// left value - right value then add the total
	var total int
	for i, val := range left {
		var distance int
		rVal := right[i]
		if val > rVal {
			distance = val - rVal
		} else {
			distance = rVal - val
		}
		total += distance
	}
	return total
}
