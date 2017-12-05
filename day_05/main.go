package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func processInstructionsPartOne(arr []int) int {
	instructions := make([]int, len(arr))
	copy(instructions, arr)
	jumps := 0
	escaped := false

	index := 0
	for !escaped {
		offset := instructions[index]
		instructions[index]++
		index += offset
		jumps++
		if index >= len(instructions) {
			escaped = true
		}
	}

	return jumps
}

func processInstructionsPartTwo(arr []int) int {
	instructions := make([]int, len(arr))
	copy(instructions, arr)
	jumps := 0
	escaped := false

	index := 0
	for !escaped {
		offset := instructions[index]
		if offset >= 3 {
			instructions[index]--
		} else {
			instructions[index]++
		}
		index += offset
		jumps++
		if index >= len(instructions) {
			escaped = true
		}
	}

	return jumps
}

func convertToInts(b []byte) []int {
	strList := strings.Split(string(b), "\n")
	intList := make([]int, len(strList))
	for i, str := range strList {
		intList[i], _ = strconv.Atoi(str)
	}
	return intList
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename", os.Args[0])
		return 1
	}

	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	arr := convertToInts(f)

	answer1 := processInstructionsPartOne(arr)
	answer2 := processInstructionsPartTwo(arr)

	fmt.Println("Answer one: ", answer1)
	fmt.Println("Answer two: ", answer2)

	return 0
}

func main() {
	os.Exit(run())
}
