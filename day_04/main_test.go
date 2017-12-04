package main

import (
	"testing"
	"fmt"
)

func TestIsAnagramDifferentSizedWords(t *testing.T) {
	input := []string{"a", "aa"}
	expected := false
	actual := isAnagramOf(input[0], input[1])
	if actual != expected {
		fmt.Errorf("isAnagramOf(%s, %s): expected %t, got %t", input[0], input[1], expected, actual)
	}
}

func TestIsAnagramOfWithSingleA(t *testing.T) {
	input := []string{"a", "a"}
	expected := true
	actual := isAnagramOf(input[0], input[1])
	if actual != expected {
		fmt.Errorf("isAnagramOf(%s, %s): expected %t, got %t", input[0], input[1], expected, actual)
	}
}

func TestIsAnagramOfWithDoubleA(t *testing.T) {
	input := []string{"aa", "aa"}
	expected := true
	actual := isAnagramOf(input[0], input[1])
	if actual != expected {
		fmt.Errorf("isAnagramOf(%s, %s): expected %t, got %t", input[0], input[1], expected, actual)
	}
}

func TestIsAnagramOfWithOneEmpty(t *testing.T) {
	input := []string{"aa", ""}
	expected := false
	actual := isAnagramOf(input[0], input[1])
	if actual != expected {
		fmt.Errorf("isAnagramOf(%s, %s): expected %t, got %t", input[0], input[1], expected, actual)
	}
}

func TestIsAnagramOfReversedWords(t *testing.T) {
	input := []string{"abcdef", "fedcba"}
	expected := true
	actual := isAnagramOf(input[0], input[1])
	if actual != expected {
		fmt.Errorf("isAnagramOf(%s, %s): expected %t, got %t", input[0], input[1], expected, actual)
	}
}
