package main

import (
	"log"
	"os"
	"testing"
)

func TestValidCharOcc(t *testing.T) {
	input := "./135-0.txt"
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}

	resultMap := countChars(f)
	expectedX, expectedt := resultMap["X"], resultMap["t"]

	if expectedX != 333 {
		t.Errorf("Expected 333 occurrences of 'X', got %v instead.", expectedX)
	}

	if expectedt != 223000 {
		t.Errorf("Expected 223000 occurrences of 't', got %v instead.", expectedt)
	}
}