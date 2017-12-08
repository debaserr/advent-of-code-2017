package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	os.Exit(run())
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage %s filename\n", os.Args[0])
		return 1
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	defer f.Close()

	registerName, registerValue, highestSeen := process(f)
	fmt.Printf("Largest register value is %d in register %s\nThe highest seen value is %d\n", registerValue, registerName, highestSeen)

	return 0
}

func process(f io.Reader) (string, int, int) {
	scanner := bufio.NewScanner(f)
	registers := make(map[string]int)
	highestSeen := 0
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, " if ")
		if parseAndEvaluateBooleanExpression(parts[1], registers) {
			val := parseAndCalculateArithmeticExpression(parts[0], registers)
			if val > highestSeen {
				highestSeen = val
			}
		}
	}

	maxVal := 0
	var maxValKey string
	for k, v := range registers {
		fmt.Println(k, ": ", v)
		if maxValKey == "" || v > maxVal {
			maxValKey = k
			maxVal = v
		}
	}

	return maxValKey, maxVal, highestSeen
}

// Parse and evaluate a boolean expression formatted as a > 1
func parseAndEvaluateBooleanExpression(ex string, registers map[string]int) bool {
	registerName, operator, right := parseThreePartExpression(ex)
	registerVal := getValueFromMap(registerName, registers)
	return evalBool(registerVal, right, operator)
}

// Parse and calculate an arithmetic expression formatted as b inc 5
func parseAndCalculateArithmeticExpression(ex string, registers map[string]int) int {
	registerName, operator, right := parseThreePartExpression(ex)
	registerVal := getValueFromMap(registerName, registers)
	newVal := evalArithmetic(registerVal, right, operator)
	registers[registerName] = newVal
	return newVal
}

func evalBool(left, right int, operator string) bool {
	switch operator {
	case "==":
		return left == right
	case "!=":
		return left != right
	case ">":
		return left > right
	case "<":
		return left < right
	case ">=":
		return left >= right
	case "<=":
		return left <= right
	default:
		fmt.Println("WARNING: evalBool ", operator, " return false!")
		return false
	}
}

func evalArithmetic(left, right int, operator string) int {
	switch operator {
	case "inc":
		return left + right
	case "dec":
		return left - right
	default:
		return 0
	}
}

// Parse three part expression formatted as x op int
func parseThreePartExpression(ex string) (string, string, int) {
	parts := strings.Split(ex, " ")
	left := parts[0]
	operator := parts[1]
	right, _ := strconv.Atoi(parts[2])
	return left, operator, right
}

func getValueFromMap(key string, m map[string]int) int {
	v, ok := m[key]
	if !ok {
		m[key] = 0
	}
	return v
}
