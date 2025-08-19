package main

import (
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	file, err := os.ReadFile("../inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(string(file))
}

func TestPart2(t *testing.T) {
	file, err := os.ReadFile("../inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}

	part2(string(file))

}
