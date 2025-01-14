// --- Day 5: If You Give A Seed A Fertilizer ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type mapType struct {
	title  string
	values [][]int
}

func part1(file *os.File) {
	start := time.Now()
	scanner := bufio.NewScanner(file)
	var seeds []int
	var maps []mapType

	var title string
parsing:
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, ":")

		// get the seeds
		if len(seeds) == 0 {
			seedsStr := strings.Split(strings.TrimSpace(parts[1]), " ")
			for _, seed := range seedsStr {
				num, err := strconv.Atoi(seed)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, num)
			}
			continue parsing
		}

		// skip blanks
		if text == "" {
			title = ""
			continue parsing
		}

		// get the title
		if strings.Contains(text, "map") {
			title = strings.Split(text, " ")[0]
			newMap := mapType{title: title}
			maps = append(maps, newMap)
			continue parsing
		}

		var nums []int
		numStr := strings.Split(text, " ")
		for _, str := range numStr {
			num, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
	insert:
		for i, mapList := range maps {
			if mapList.title == title {
				maps[i].values = append(mapList.values, nums)
				break insert
			}
		}
	}

	var location int
search:
	for _, seed := range seeds {
		target := seed

		for _, mapList := range maps {
			numArrs := mapList.values
		numRange:
			for _, numArr := range numArrs {
				desNum := numArr[0]
				srcNum := numArr[1]
				length := numArr[2]

				// search the map
				for j := 0; j < length; j++ {
					if target == srcNum {
						target = desNum + j
						break numRange
					}
					srcNum++
				}
			}
		}

		fmt.Println("Seed done:", seed)

		if location == 0 {
			location = target
			continue search
		}

		if target < location {
			location = target
		}

	}

	fmt.Println("location:", location)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	part1(file)
}
