package main
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


func TestPositifNumber(t *testing.T){
	input := []Frame{{0,1},{13,2},{0,1},{13,2},{0,1},{13,2},{0,1},{13,2},{0,1},{13,2}}
	expected := 0
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}


func TestScore(t *testing.T) {
	input := []Frame{{1,2},{1,2},{1,2},{1,2},{1,2},{1,2},{1,2},{1,2},{1,2},{1,2}}
	expected := 30
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}
	


func TestNumberOfIndex(t *testing.T){
	input := []Frame{{4,5},{5,6},{4,5},{5,6},{4,5},{5,6},{4,5},{5,6},{4,5},{5,6}}
	expected := 200
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}


func TestSomme(t *testing.T) {
	input := []Frame{{10, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}}
	expected := 0
	if err := scoreChecker(input, expected, errExpected); err != nil {
		t.Fatalf("%+v\n", err)
	}
}


func testspare(t *testing.T) { 
    input := []Frame{{8,2}, {5,1},{0,0},{0,0},{0,0},{0,0},{0,0},{0,0},{0,0},{0,0}}
	expected := 10+5+5+1
	
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestSpareEnd(t *testing.T) {
	input := []Frame{{0, 3}, {5, 1}, {3, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {6,4}}
	expected := 0+3+5+1+3+6+4

	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestStrike(t *testing.T) {
	input := []Frame{{10, 0}, {6,2}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expected := 10+6+2+6+2

	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestStrikeEnd(t *testing.T) {
	input := []Frame{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {10, 0}}
	expected := 10
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

