package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(file *os.File) {
	// max of R, G, B cubes according to "https://adventofcode.com/2023/day/2"
	const RED, GREEN, BLUE = 12, 13, 14
	var total int
	scanner := bufio.NewScanner(file)

text:
	for scanner.Scan() {
		text := scanner.Text()
		// split the input to two (game number and cubes shown)
		// e.g. "Game 1: 3 b, 4 r..." to ["Game 1:", "3 b, 4 r..."]
		parts := strings.Split(text, ":")

		// split the cubes shown sets
		// e.g. "4 r; 2 g" to ["4 r", "2 g"]
		game := parts[1]
		set := strings.Split(game, ";")

		for _, subset := range set {
			// in a subset separated by ";"
			// a cube should not exceed the maximum
			// number of cubes for that particular color
			r, g, b := 0, 0, 0
			// split the subset of cubes
			cubes := strings.Split(subset, ",")

			for _, cube := range cubes {
				// get each cube number and color
				cube = strings.TrimSpace(cube)
				cubeDetails := strings.Split(cube, " ")
				cubeNum := cubeDetails[0]
				cubeColor := cubeDetails[1]

				num, err := strconv.Atoi(cubeNum)
				if err != nil {
					panic("Unable to convert string to num")
				}

				switch cubeColor {
				case "red":
					r += num
				case "green":
					g += num
				case "blue":
					b += num
				}

				if r > RED || g > GREEN || b > BLUE {
					// r, g, b exceed the max number of cubes per color
					// this game is invalid
					// proceed to next game
					continue text
				}

			}
		}

		// if we get into this point this game is valid
		// all we have to do is split to get the id and
		// add it all up to get the answer
		// e.g "Game 1" to ["Game", "1"]
		gameDetails := parts[0]
		gameDetSplit := strings.Split(gameDetails, " ")
		gameId := gameDetSplit[1]
		id, err := strconv.Atoi(gameId)
		if err != nil {
			panic("Unable to convert game id to number")
		}
		total += id

	}

	fmt.Println("Part 1:", total)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Unable to open file")
	}
	defer file.Close()

	part1(file)
}
