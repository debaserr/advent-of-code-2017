package main

import (
	"testing"
	"fmt"
)

func TestAddOne(t *testing.T) {
	expectedFirstVal := 1
	g := GetGrid()
	g.Add()
	currVal := g.GetCurrentValue()
	if currVal != expectedFirstVal {
		t.Errorf("currVal: %d, expected: %d", currVal, expectedFirstVal)
	}
}

func TestAddTwo(t *testing.T) {
	expected := 2
	g := GetGrid()
	g.Add()
	g.Add()
	currVal := g.GetCurrentValue()
	if currVal != expected {
		t.Errorf("currVal: %d, expected: %d", currVal, expected)
	}
}

func TestAddThree(t *testing.T) {
	expected := 4
	g := GetGrid()
	g.Add()
	g.Add()
	g.Add()
	currVal := g.GetCurrentValue()
	if currVal != expected {
		t.Errorf("currVal: %d, expected: %d", currVal, expected)
	}
}

func TestFullCircle(t *testing.T) {
	g := GetGrid()
	steps := []int{1, 2, 4, 5, 10, 11, 23, 25, 26}
	for _, expected := range steps {
		g.Add()
		currVal := g.GetCurrentValue()
		if currVal != expected {
			t.Errorf("currVal: %d, expected: %d", currVal, expected)
		}
	}
}

func TestPuzzleAnswer(t *testing.T) {
	g := GetGrid()
	gridVal := 0
	input := 265149
	for gridVal < input {
		g.Add()
		gridVal = g.GetCurrentValue()
	}
	fmt.Println("Answer is ", gridVal)
}
