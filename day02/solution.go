package main

import (
	"fmt"
	"log"
	"os"
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

func parseInput(input string) [][]int {
	var reports [][]int

	for _, line := range strings.Split(input, "\n") {
		var parts = strings.Fields(line)
		var report []int

		for _, part := range parts {
			var i, err = strconv.Atoi(part)
			if err != nil {
				log.Fatal("Unable to parse part", part)
			}
			report = append(report, i)
		}

		reports = append(reports, report)
	}

	return reports
}

func part1() {
	var input = readInput("./inputs/day02.txt")
	var reports = parseInput(input)

	var totalSafe = 0
reportLoop:
	for _, report := range reports {
		var increase = (report[1] - report[0]) < 0
		for i := 1; i < len(report); i++ {
			var diff = report[i] - report[i-1]
			if abs(diff) <= 0 || abs(diff) > 3 {
				continue reportLoop
			}

			if increase != (diff < 0) {
				continue reportLoop
			}

		}

		totalSafe += 1
	}

	fmt.Println("Part 1: ", totalSafe)
}

func part2() {
	var input = readInput("./inputs/day02.txt")
	var reports = parseInput(input)

	var totalSafe = 0
reportLoop:
	for _, report := range reports {
		if isSafe(report) {
			totalSafe += 1
			continue
		}

		for i := 0; i < len(report); i++ {
			var tmp = append([]int(nil), report[:i]...)
			tmp = append(tmp, report[i+1:]...)
			if isSafe(tmp) {
				totalSafe += 1
				continue reportLoop
			}
		}
	}

	fmt.Println("Part 2: ", totalSafe)
}

func isSafe(report []int) bool {
	var increase = (report[1] - report[0]) < 0
	for i := 1; i < len(report); i++ {
		var diff = report[i] - report[i-1]
		if abs(diff) <= 0 || abs(diff) > 3 {
			return false
		}

		if increase != (diff < 0) {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
