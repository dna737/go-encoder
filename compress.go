package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type HuffNode struct {
	weight int
	left string
	right string
	isLeaf bool
}

func countChars(f *os.File) map[string]int {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)
	occ := make(map[string]int)

	for scanner.Scan() {
		a := string(scanner.Bytes())
		if unicode.IsLetter(rune(a[0])) {
			if val, ok := occ[a]; ok {
				occ[a] = val + 1
			} else {
				occ[a] = 1
			}
		}
	}

	return occ
}

func compGen(occ map[string]int) ([]string, []HuffNode) {
	occCopy := occ
	var result []string
	var nodes []HuffNode

	for len(occCopy) > 0 {
		n, k1, k2 := generateNodes(occCopy)
		fmt.Println(k1, k2)
		delete(occCopy, k1)
		delete(occCopy, k2)

		result = append(result, k1, k2)
		nodes = append(nodes, n)
	}

	return result, nodes
}

func buildTree(nodes []HuffNode) HuffNode {
	var result HuffNode

	//Nodes are in increasing order of freq, and will form pairs first.
	for len(nodes) != 1 {
		for i, n := range nodes {
			var secondVal int

			if i + 1 == len(nodes) {
				//No second node
				secondVal = 0
			} else {
				secondVal = nodes[i + 1].weight
			}

			newNode := HuffNode{
				weight: n.weight + secondVal,
				left: "",
				right: "",
				isLeaf: false,
			}

			// nodes
		}
	}
}

func generateNodes(occ map[string]int) (HuffNode, string, string) {

	smallest, secondSmallest := 0, 0
	var smallestChar, secondSmallestChar string
	
	for k, v := range occ {

		if smallest == 0 {
			smallest = v
			smallestChar = k
		} else if secondSmallest == 0 {
			secondSmallest = v
			secondSmallestChar = k
			if v < smallest {
				secondSmallest = smallest
				secondSmallestChar = smallestChar
				smallest = v
				smallestChar = k
				continue
			}
		} else {
			if v < smallest {
				secondSmallest = smallest
				secondSmallestChar = smallestChar
				smallest = v
				smallestChar = k
				continue
			} else if v < secondSmallest {
				secondSmallest = v
				secondSmallestChar = k
			}
		}
	}

	a := HuffNode{
		weight: (smallest + secondSmallest),
		left:   smallestChar,
		right:  secondSmallestChar,
		isLeaf: true,
	}

	fmt.Println(a)

	return a, smallestChar, secondSmallestChar
}

func main(){

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Please provide a valid path to the file name as an option.")
	}
	
	countChars(f)
}