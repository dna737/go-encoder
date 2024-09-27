package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type HuffNode struct {
	weight int
	value string
	left *HuffNode
	right *HuffNode
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

func compGen(occ map[string]int) ([]HuffNode) {
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

	return nodes
}

func buildTree(nodes []HuffNode) HuffNode {

	//Nodes are in increasing order of freq, and will form pairs first.
	for len(nodes) != 1 {

		for i := 0; i < len(nodes); i++ {
			var secondVal int

			if i + 1 == len(nodes) {
				continue //Single nodes are left alone.
			} else {
				secondVal = nodes[i + 1].weight
			}

			newNode := HuffNode{
				weight: nodes[i].weight + secondVal,
				left: &nodes[i],
				right: &nodes[i + 1],
				isLeaf: false,
			}
			nodes = append(append(nodes[:i],nodes[i + 2:]...), newNode)
			fmt.Println(nodes)
		}
	}

	return nodes[0]
}

//Generates a node from the two smallest values in the map.
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
		left:   &HuffNode{value: smallestChar},
		right:   &HuffNode{value: secondSmallestChar},
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
