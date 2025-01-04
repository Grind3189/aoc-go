// day2
// Analyze the data whether its safe or not base on the 2 condition
// 1.The level are either increasing or decreasing
// 2. Differ by at least 1 or at most 3
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var report int
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open: %v\n", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		parts := strings.Fields(scanner.Text())
		passed := perfCheck(parts)

		if passed == false {
			// remove a single element and perform the check again
			for i := range parts {
				newArr := removeIndex(parts, i)
				passed = perfCheck(newArr)
				if passed {
					break
				}
			}
		}

		if passed {
			report++
		}

	}

	fmt.Println("Report:", report)
}

func removeIndex(arr []string, i int) []string {
	temp := make([]string, 0)
	// append the elemets before the index provided to the new arr
	temp = append(temp, arr[:i]...)

	// append the elements after the index provided and return the new array
	return append(temp, arr[i+1:]...)
}

func perfCheck(arr []string) bool {
	var increasing bool
	var decreasing bool
	var passed bool
	// loop through each element of a level
	for i, num := range arr {
		// Im on the last index and it should be settled on the previous loop
		if i+1 == len(arr) {
			passed = true
			break
		}

		num, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error converting string to number")
			break
		}

		num2, err := strconv.Atoi(arr[i+1])
		if err != nil {
			fmt.Println("Error converting string to number")
			break
		}

		if num < num2 { // increasing
			increasing = true
			diff := num2 - num
			// levels should be increasing or decreasing only not both
			if decreasing || diff > 3 || diff < 1 {
				break
			}

		} else if num > num2 { // decreasing
			decreasing = true
			diff := num - num2

			if increasing || diff > 3 || diff < 1 {
				break
			}

		} else {
			// neither increasing or decreasing
			break
		}
	}

	return passed
}
