package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func getChecksumDiv(input string) int {
	sum := 0
	rows := strings.Split(input, "\n")

	for _, row := range rows {
		blocks := strings.Split(row, "\t")
		nums := make([]int, len(blocks))
		for i, block := range blocks {
			num, _ := strconv.Atoi(block)
			for j := 0; j < i; j++ {
				prevNum := nums[j]
				if num % prevNum == 0 {
					sum += num / prevNum
					continue
				}
				if prevNum % num == 0 {
					sum += prevNum / num
				}
			}
			nums[i] = num
		}
	}

	return sum
}

func getChecksum(input string) int {
	sum := 0
	rows := strings.Split(input, "\n")

	for _, row := range rows {
		blocks := strings.Split(row, "\t")
		var minVal, maxVal int
		minValSet, maxValSet := false, false
		for _, block := range blocks {
			num, _ := strconv.Atoi(block)

			if !maxValSet || num > maxVal {
				maxVal = num
				maxValSet = true
			}

			if !minValSet || num < minVal {
				minVal = num
				minValSet = true
			}
		}
		sum += maxVal - minVal
	}

	return sum
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

	checksum := getChecksum(string(fileData))
	fmt.Printf("Checksum: %d\n", checksum)

	checksumDiv := getChecksumDiv(string(fileData))
	fmt.Printf("ChecksumDiv: %d\n", checksumDiv)

	return 0
}

func main() {
	os.Exit(run())
}
