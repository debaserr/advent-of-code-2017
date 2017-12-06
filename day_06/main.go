package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
)

func main() {
	os.Exit(run())
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
		return 1
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	banks := getBlocksFromInput(input)
	cycles := reallocateBanks(banks)

	fmt.Printf("Task 1: %d cycles done\n", cycles)

	cycles = reallocateBanks(banks)

	fmt.Printf("Task 2: %d cycles done\n", cycles)

	return 0
}

func getBlocksFromInput(input []byte) []int {
	inputArr := strings.Split(string(input), "\t")
	ret := make([]int, len(inputArr))
	for i, s := range inputArr {
		ret[i], _ = strconv.Atoi(s)
	}
	return ret
}

func reallocateBanks(banks []int) int {
	cycles := 0
	previousConfs := []string{getConfStrFromBanks(banks)}
	confRepeated := false

	for !confRepeated {
		bankIndex := getIndexOfBankWithMostBlocks(banks)
		redistributeBlocks(banks, bankIndex)
		confStr := getConfStrFromBanks(banks)
		if contains(previousConfs, confStr) {
			confRepeated = true
		} else {
			previousConfs = append(previousConfs, confStr)
		}
		cycles++
	}

	return cycles
}

func redistributeBlocks(banks []int, indexToEmpty int) {
	blocksToRedistribute := banks[indexToEmpty]
	banks[indexToEmpty] = 0
	i := indexToEmpty + 1
	for blocksToRedistribute > 0 {
		if i >= len(banks) {
			i = 0
		}
		banks[i]++
		blocksToRedistribute--
		i++
	}
}

func getConfStrFromBanks(banks []int) string {
	ret := ""
	for _, bank := range banks {
		ret += fmt.Sprintf("%d ", bank)
	}
	return strings.TrimRight(ret, " ")
}

func getIndexOfBankWithMostBlocks(banks []int) int {
	high := -1
	var bankIndexWithMostBlocks int
	for i, bank := range banks {
		if bank > high {
			high = bank
			bankIndexWithMostBlocks = i
		}
	}

	return bankIndexWithMostBlocks
}

func contains(arr []string, e string) bool {
	for _, s := range arr {
		if e == s {
			return true
		}
	}
	return false
}
