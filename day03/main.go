package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("inputs/day03.txt")
	fmt.Println("Part 1:", part1(string(data)))
	fmt.Println("Part 2:", part2(string(data)))
}
