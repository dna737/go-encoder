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

// func TestValidHuffmanTree(t *testing.T){
// 	occ := make(map[string]int)
// 	occ["A"] = 3
// 	occ["B"] = 2
// 	occ["C"] = 1
// 	tree, ok := generateBinaryTree(occ)
// 	if !ok{
// 		t.Errorf("wrong type.")
// 	}
// 	fmt.Println(tree.weight)
// }

func TestCompGen(t *testing.T){
	occ := make(map[string]int)
	occ["a"] = 1
	occ["i"] = 2
	occ["t"] = 3
	occ["e"] = 1
	occ["s"] = 3
	occ["h"] = 1

	nodeSlice := compGen(occ)
	fmt.Println(nodeSlice)
}

func TestBuildTree(t *testing.T){
	occ := make(map[string]int)
	occ["a"] = 1
	occ["i"] = 2
	occ["t"] = 3
	occ["e"] = 1
	occ["s"] = 3
	occ["h"] = 1

	nodeSlice := compGen(occ)
	table := make(map[string]string)
	finalNode := buildTree(nodeSlice)
	result := printValues(finalNode, "", table)
	fmt.Println(result)
}
