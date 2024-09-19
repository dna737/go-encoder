package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type HuffTree struct {
	weight int
	left string
	right string
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

func compGen(occ map[string]int) []string {
	occCopy := occ
	var result []string

	for len(occCopy) > 0 {
		_, k1, k2 := generateNodes(occCopy)
		fmt.Println(k1, k2)
		delete(occCopy, k1)
		delete(occCopy, k2)

		result = append(result, k1, k2)
	}

	return result
}

func generateNodes(occ map[string]int) (HuffTree, string, string) {

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

	a := HuffTree{
		weight: (smallest + secondSmallest),
		left:   smallestChar,
		right:  secondSmallestChar,
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