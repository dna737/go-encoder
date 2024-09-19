package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func openFile(filepath string) *os.File {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}

	return f
}

func TestValidCharOcc(t *testing.T) {
	input := "./135-0.txt"
	f := openFile(input)

	resultMap := countChars(f)
	expectedX, expectedt := resultMap["X"], resultMap["t"]

	if expectedX != 333 {
		t.Errorf("Expected 333 occurrences of 'X', got %v instead.", expectedX)
	}

	if expectedt != 223000 {
		t.Errorf("Expected 223000 occurrences of 't', got %v instead.", expectedt)
	}
}

func TestValidHuffmanTree(t *testing.T){
	occ := make(map[string]int)
	occ["A"] = 3
	occ["B"] = 2
	occ["C"] = 1
	res1, res2 := generateBinaryTree(occ)

	if res1 != 1 || res2 != 2 {
		fmt.Println(res1, res2)
		t.Errorf("wrong!")
	}

}