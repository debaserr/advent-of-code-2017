package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"strings"
	"strconv"
)

type ProgramCorrection struct{
	program *Program
	weightAdjustment int
}

type CalculatedChild struct {
	p *Program
	calculatedWeight int
}

func createTowerFromInput(f io.Reader) *Program {
	scanner := bufio.NewScanner(f)
	var parents []*Program
	var orphans []*Program
	for scanner.Scan() {
		s := scanner.Text()
		name, weight, children := extractData(s)
		program := NewProgram(name, weight, children)
		if program.IsParent() {
			parents = append(parents, program)
		} else {
			orphans = append(orphans, program)
		}
	}

	for _, parent := range parents {
		for childIndex, child := range parent.children {
			foundChild := false
			for orphanIndex, orphan := range orphans {
				if child.name == orphan.name {
					orphan.parent = parent
					parent.children[childIndex] = orphan
					orphans = append(orphans[:orphanIndex], orphans[orphanIndex + 1:]...)
					foundChild = true
					break
				}
			}
			if !foundChild {
				for _, p := range parents {
					if child.name == p.name {
						p.parent = parent
						parent.children[childIndex] = p
						break
					}
				}
			}
		}
	}

	return findRoot(parents)
}

func findRoot(programs []*Program) *Program {
	for _, p := range programs {
		if p.parent == nil {
			return p
		}
	}
	return nil
}

func getWeight(p *Program) (int, *ProgramCorrection) {
	children := make([]CalculatedChild, len(p.children))
	childrenWeight := 0
	var correction *ProgramCorrection
	for i, child := range p.children {
		childWeight, corr := getWeight(child)
		children[i] = CalculatedChild{child, childWeight}
		childrenWeight += childWeight
		if corr != nil {
			correction = corr
		}
	}

	if correction == nil && len(p.children) > 0 {
		weightDiff, unevenChild := findUneven(children)
		if weightDiff != 0 {
			correction = &ProgramCorrection{
				unevenChild.p,
				weightDiff,
			}
		}
	}

	weight := p.weight + childrenWeight
	return weight, correction
}

func findUneven(children []CalculatedChild) (int, CalculatedChild) {
	diff := 0
	var unevenChild CalculatedChild

	m := make(map[int][]CalculatedChild)
	for _, child := range children {
		m[child.calculatedWeight] = append(m[child.calculatedWeight], child)
	}

	if len(m) > 1 {
		var unevenValue, correctValue int
		for k, v := range m {
			if len(v) > 1 {
				correctValue = k
			} else {
				unevenValue = k
				unevenChild = m[k][0]
			}
		}
		diff = correctValue - unevenValue
	}

	return diff, unevenChild
}

func extractData(s string) (string, int, []string) {
	var children []string
	split := strings.Split(s, " -> ")
	if len(split) > 1 {
		children = strings.Split(split[1], ", ")
	}
	nameWeight := strings.Split(split[0], " ")
	name := nameWeight[0]
	weightStr := strings.Replace(nameWeight[1], "(", "", 1)
	weightStr = strings.Replace(weightStr, ")", "", 1)
	weight,_ := strconv.Atoi(weightStr)
	return name, weight, children
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
		return 1
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	defer f.Close()

	tower := createTowerFromInput(f)

	fmt.Println("Bottom program is: ", tower.name)

	weight, corr := getWeight(tower)
	fmt.Printf("Total weight: %d\nUneven program: %s\nCorrection: %d\nNew weight: %d\n",
		weight, corr.program.name, corr.weightAdjustment, corr.program.weight + corr.weightAdjustment)

	return 0
}

func main() {
	os.Exit(run())
}
