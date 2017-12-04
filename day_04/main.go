package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"strings"
)

func isAnagramOf(subject string, candidate string) bool {
	if len(subject) != len(candidate) {
		return false
	}
	for _, r := range subject {
		if strings.ContainsRune(candidate, r) {
			candidate = strings.Replace(candidate, string(r), "", 1)
		} else {
			return false
		}
	}
	return true
}

func getAnswers(f io.Reader) (int, int) {
	numValidSame := 0
	numValidAnagrams := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		containsDuplicates := false
		containsAnagrams := false

		s := scanner.Text()
		words := strings.Split(s, " ")
		for i, word := range words {
			for _, otherWord := range words[i + 1:] {
				if word == otherWord {
					containsDuplicates = true
				}
				if isAnagramOf(word, otherWord) {
					containsAnagrams = true
				}
			}
		}

		if !containsDuplicates {
			numValidSame++
		}
		if !containsAnagrams {
			numValidAnagrams++
		}
	}

	return numValidSame, numValidAnagrams
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename", os.Args[0])
		return 1
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	defer f.Close()

	answer1, answer2 := getAnswers(f)
	fmt.Println("Answer one is ", answer1)
	fmt.Println("Answer two is ", answer2)

	return 0
}

func main() {
	os.Exit(run())
}
