package main

import (
	"testing"
)

func TestOneEmptyGroup(t *testing.T) {
	input := "{}"
	expected := 1
	actual, _ := GetScore(input)
	if actual !=expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestTwoEmptyGroups(t *testing.T) {
	input := "{{}}"
	expected := 3
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestThreeEmptyGroups(t *testing.T) {
	input := "{{{}}}"
	expected := 6
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestTenEmptyGroups(t *testing.T) {
	input := "{{{{{{{{{{}}}}}}}}}}"
	expected := 55
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestTwoInnerGroups(t *testing.T) {
	input := "{{},{}}"
	expected := 5
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestTwoInnerWithOneInnerGroups(t *testing.T) {
	input := "{{},{{}}}"
	expected := 8
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestThreeInnerGroups(t *testing.T) {
	input := "{{{},{},{{}}}}"
	expected := 16
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestFourGarbagesSeparatedByCommas(t *testing.T) {
	input := "{<a>,<a>,<a>,<a>}"
	expected := 1
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestGarbageWithInnerGroup(t *testing.T) {
	input := "{<{}>}"
	expected := 1
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestGarbageWithMultipleInnerGroups(t *testing.T) {
	input := "{<{},{},{{}}>}"
	expected := 1
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestFourInnerGroupsWithOneInnerGarbageEach(t *testing.T) {
	input := "{{<ab>},{<ab>},{<ab>},{<ab>}}"
	expected := 9
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestGarabgeWithIgnoredEnding(t *testing.T) {
	input := "<!>{}>"
	expected := 0
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestGarabgeWithLotsOfGarbage(t *testing.T) {
	input := "<{o\"i!a,<{i<a>"
	expected := 0
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestGarabgeWithOuterGroup(t *testing.T) {
	input := "{<{o\"i!a,<{i<a>}"
	expected := 1
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestTwoGroupsWithGarbageAndCancelations(t *testing.T) {
	input := "{{<!>},{<!>},{<!>},{<a>}}"
	expected := 3
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}

func TestInnerGroupsWithGarbageAndDoubleCancelations(t *testing.T) {
	input := "{{<!!>},{<!!>},{<!!>},{<!!>}}"
	expected := 9
	actual, _ := GetScore(input)
	if actual != expected {
		t.Errorf("GetScore: expected %d, got %d", expected, actual)
	}
}
