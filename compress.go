package main

import (
	"bufio"
	"fmt"
	"os"
)

func countChars(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)
	arr := []int{}

	for scanner.Scan() {
		// append(arr, )
	}
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
