package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput(name string) string {
	file, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}

func parseInput(input string) map[string][]int {
	var pairs = map[string][]int{
		"left":  {},
		"right": {},
	}

	for _, line := range strings.Split(input, "\n") {
		var parts = strings.Fields(line)

		var left, err1 = strconv.Atoi(parts[0])
		var right, err2 = strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			log.Fatal("Unable to parse line", line)
		}

		pairs["left"] = append(pairs["left"], left)
		pairs["right"] = append(pairs["right"], right)
	}

	slices.Sort(pairs["left"])
	slices.Sort(pairs["right"])

	return pairs
}

func part1(input string) int {
	var pairs = parseInput(input)

	var totalDistance = 0
	for i, _ := range pairs["left"] {
		totalDistance += abs(pairs["left"][i] - pairs["right"][i])
	}

	return totalDistance
}

func part2(input string) int {
	var pairs = parseInput(input)

	var totalSimilarity = 0
	for _, left := range pairs["left"] {
		var count = 0
		for _, right := range pairs["right"] {
			if left == right {
				count++
			}
		}

		totalSimilarity += count * left
	}

	return totalSimilarity
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
