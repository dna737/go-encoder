package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func countChars(f *os.File) map[string]int {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)
	occ := map[string]int{}

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


func main() {

	if len(os.Args[1:]) != 1 {
		fmt.Println("Please provide a valid path to the file name as an option.")
		os.Exit(1)
	}
	
	filename := os.Args[1]
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("Please provide a valid path to the file name as an option.")
	}
	
	countChars(f)
	
}
