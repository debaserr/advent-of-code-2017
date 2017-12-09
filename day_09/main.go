package main

import (
	"os"
	"fmt"
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

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	score, garbageChars := GetScore(string(b))

	fmt.Println("The score is ", score)
	fmt.Println("Number of garbage characters are ", garbageChars)

	return 0
}

func GetScore(in string) (int, int) {
	const (
		GroupStart = "{"
		GroupEnd = "}"
		GarbageStart = "<"
		GarbageEnd = ">"
		Ignore = "!"
	)
	score, level, garbageChars := 0, 0, 0
	inGarbage, ignoreNext := false, false
	for _, r := range in {
		if ignoreNext {
			ignoreNext = false
			continue
		}

		char := string(r)

		if char == Ignore {
			ignoreNext = true
			continue
		}

		if inGarbage {
			if char == GarbageEnd {
				inGarbage = false
			} else {
				garbageChars++
			}
			continue
		}

		switch char {
		case GroupStart:
			level++
			break
		case GroupEnd:
			score += level
			level--
		case GarbageStart:
			inGarbage = true
		}
	}
	return score, garbageChars
}
