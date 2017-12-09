package main

import (
	"testing"
)

func TestAnswerTwo(t *testing.T) {
	expected := 10
	input := []int{0, 3, 0, 1, -3}
	actual := processInstructionsPartTwo(input)
	if actual != expected {
		t.Errorf("processInstructionsPartTwo: expected %d, got %d", expected, actual)
	}
}
