package bowling

import (
	"fmt"
	"testing"
)

func scoreChecker(input []Frame, expectedScore int, expectedError error) error {
	score, err := GetScore(input)

	if err != expectedError && !(err != nil && expectedError != nil && err.Error() == expectedError.Error()) {
		return fmt.Errorf("Score error : %+v, expected %+v", err, expectedError)
	}
	if score != expectedScore {
		return fmt.Errorf("Score : %d, expected %d", score, expectedScore)
	}
	return nil
}

func TestNullScore(t *testing.T) {
	input := []Frame{}
	expected := 0
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScore(t *testing.T) {
	input := []Frame{{0,1},{1,2},{0,1},{1,2},{0,1},{1,2},{0,1},{1,2},{0,1},{1,2}}
	expected := 20
	
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
	
}

func TestNumberOfIndex(t *testing.T){
	input := []Frame{{0,1},{1,2},{0,1},{1,2},{0,1},{1,2},{0,1},{1,2},{0,1},{1,2}}
	expected := 20
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestNumberPositif(t *testing.T){
	input := []Frame{{0,1},{13,2},{0,-1},{13,2},{0,1},{13,2},{0,1},{13,2},{0,1},{13,2}}
	expected := 0
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestSommeTuplet(t *testing.T){
	input := []Frame{{0,1},{13,2},{0,1},{13,2},{0,1},{13,2},{0,1},{13,2},{0,1},{13,2}}
	expected := 0
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestSpare(t *testing.T) {
	input := []Frame{{7, 3}, {5, 1}, {3, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expected := 7+3+5+5+1+3

	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestSpareAtTheEnd(t *testing.T) {
	input := []Frame{{0, 3}, {5, 1}, {3, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {7,3}}
	expected := 0+3+5+1+3+7+3

	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestStrike(t *testing.T) {
	input := []Frame{{10, 0}, {5, 3}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expected := 10+5+3+5+3

	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestStrikeAtTheEnd(t *testing.T) {
	input := []Frame{{0, 0}, {6, 1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {10, 0}}
	expected := 0+0+6+1+10

	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestDoubleStrike(t *testing.T) {
	input := []Frame{{10, 0}, {10, 0}, {5, 1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expected := 10+10+10+5+1+5+1

	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}