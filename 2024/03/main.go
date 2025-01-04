package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var total int
	do := true
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "file:%v\n", err)
	}

	fmt.Println(do)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		for i := 0; i < len(text); i++ {
			// get the current chunks from starting index to the end
			current := text[i:]
			if len(current) < 8 {
				fmt.Println("Last chunk")
				break
			}

			// get a string of 7 characters
			// exactly the same length of word don't()
			target := current[:8]
			fmt.Println(target)
			if strings.Contains(target, "don't()") {
				do = false
			}
		}
	}

	fmt.Println("total:", total)
}
