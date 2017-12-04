package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strconv"
)

func getSequenceSums(seq string) (int, int) {
	sum1 := 0
	sum2 := 0

	sequenceLength := len(seq)
	halfwayRoundOffset := sequenceLength / 2
	for i := 0; i < sequenceLength; i++ {
		nextIndex := (i + 1) % sequenceLength
		nextHalfwayIndex := (i + halfwayRoundOffset) %sequenceLength

		num, _ := strconv.Atoi(string(seq[i]))
		nextNum, _ := strconv.Atoi(string(seq[nextIndex]))
		nextHalfwayNum, _ := strconv.Atoi(string(seq[nextHalfwayIndex]))

		if num == nextNum {
			sum1 += num
		}
		if num == nextHalfwayNum {
			sum2 += num
		}
	}
	return sum1, sum2
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Syntax: %s filename\n", os.Args[0])
		return 1
	}

	fileData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	sum1, sum2 := getSequenceSums(string(fileData))
	fmt.Printf("sum1: %d, sum2: %d\n", sum1, sum2)

	return 0
}

func main() {
	os.Exit(run())
}
