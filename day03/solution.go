package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(input string) map[string][]int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	var pairs = map[string][]int{
		"x": {},
		"y": {},
	}

	for _, m := range matches {
		var x, err1 = strconv.Atoi(m[1])
		var y, err2 = strconv.Atoi(m[2])
		if err1 != nil || err2 != nil {
			log.Fatal("Unable to parse line")
		}

		pairs["x"] = append(pairs["x"], x)
		pairs["y"] = append(pairs["y"], y)
	}

	return pairs
}

type OpKind int

const (
	OpMul OpKind = iota
	OpDo
	OpDont
)

type Operation struct {
	kind   OpKind
	params []int
}

func parseInput2(input string) []Operation {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	var ops []Operation

	for _, m := range matches {
		switch {

		case strings.HasPrefix(m[0], "mul"):
			var x, err1 = strconv.Atoi(m[1])
			var y, err2 = strconv.Atoi(m[2])
			if err1 != nil || err2 != nil {
				log.Fatal("Unable to parse line")
			}

			ops = append(ops, Operation{kind: OpMul, params: []int{x, y}})

		case strings.HasPrefix(m[0], "do()"):
			ops = append(ops, Operation{kind: OpDo, params: []int{}})

		case strings.HasPrefix(m[0], "don't()"):
			ops = append(ops, Operation{kind: OpDont, params: []int{}})
		}
	}

	return ops
}

func part1(input string) int {
	var pairs = parseInput(input)
	var total = 0
	for i, _ := range pairs["x"] {
		total += pairs["x"][i] * pairs["y"][i]
	}
	return total
}

func part2(input string) int {
	var ops = parseInput2(input)

	var total = 0
	var do = true
	for i, _ := range ops {
		switch ops[i].kind {
		case OpMul:
			if !do {
				continue
			}
			total += ops[i].params[0] * ops[i].params[1]
		case OpDo:
			do = true
		case OpDont:
			do = false
		}
	}

	return total
}
