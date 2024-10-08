package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
		if val, ok := occ[a]; ok {
			occ[a] = val + 1
		} else {
			occ[a] = 1
		}
	}

	fmt.Println(occ)
	return occ
}

func compGen(occ map[string]int) ([]HuffNode) {
	occCopy := occ
	var result []string
	var nodes []HuffNode

	for len(occCopy) > 0 {
		n, k1, k2 := generateNodes(occCopy)
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
				left: &HuffNode{
					value: nodes[i].value,
					weight: nodes[i].weight,
					isLeaf: nodes[i].isLeaf,
					left: nodes[i].left,
					right: nodes[i].right,
				},
				right: &HuffNode{
					value: nodes[i + 1].value,
					weight: nodes[i + 1].weight,
					isLeaf: nodes[i + 1].isLeaf,
					left: nodes[i + 1].left,
					right: nodes[i + 1].right,
				},
				isLeaf: false,
			}

			if len(nodes) == 2 {
				nodes = append(nodes[:i], newNode)
			} else {
				nodes = append(nodes[:i],nodes[i + 2:]...)
				nodes = append(nodes, newNode)
			}
		}
	}

	return nodes[0]
}

func getPrefixTable(node HuffNode, prefix string, table map[string]string) map[string]string {
	if node.isLeaf {
		table[node.value] = prefix
	} else {
		getPrefixTable(*node.left, prefix + "0", table)
		getPrefixTable(*node.right, prefix + "1", table)
	}
	return table
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
		left:   &HuffNode{value: smallestChar, weight: smallest, isLeaf: true},
		right:   &HuffNode{value: secondSmallestChar, weight: secondSmallest, isLeaf: true},
		isLeaf: false,
	}

	return a, smallestChar, secondSmallestChar
}

func usage() {
	fmt.Println(
	`
		Usage:
			./compress-go [input-filename] [output-filename]
	`)
}

func extractBitstrings(table map[string]string, f io.Reader) []string {
	var result []string

	reader := bufio.NewReader(f)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}

		c, exists := table[string(b)]
		if !exists {
			fmt.Println("Error: Character not found in table.", string(b))
			continue
		}

		result = append(result, c)
	}

	return result	
}

func generateCompressedFile(outputFile string, table map[string]string, inputFile string) {

	file, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}

	// Starting header, prefix table, and ending header.	

	writer := bufio.NewWriter(file)
	writer.WriteString("#START#\n")
	for k, v := range table {
		writer.WriteString(k + "\t" + v + "\n")
	}
	writer.WriteString("#END#\n")

	f, err2 := os.Open(inputFile)	
	if err2 != nil {
		fmt.Println("Failed to open file (step 2).")
	}

	buffer := byte(0)
	bitCount := 0
	bitStrings := extractBitstrings(table, f)

	//Going over the bits of the current character.
    for _, bitString := range bitStrings {
        for _, bit := range bitString {
            buffer <<= 1
            if bit == '1' {
                buffer |= 1
            }
            bitCount++

            if bitCount == 8 {
                fmt.Printf("Writing byte: %08b\n", buffer) // Debugging statement
                writer.WriteByte(buffer)
                buffer = 0
                bitCount = 0
            }
        }
    }

    // Write remaining bits in the buffer
    if bitCount > 0 {
        buffer <<= (8 - bitCount)
		fmt.Printf("Writing byte: %08b\n", buffer) // Debugging statement
        writer.WriteByte(buffer)
    }

	err = writer.Flush()
	if err != nil {
	    panic(err)
	}

	defer file.Close()
}

func main(){

	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}
	
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Please provide a valid path to the file name as an option.")
	}

	table := getPrefixTable(buildTree(compGen(countChars(f))), "", make(map[string]string))
	outputFile := os.Args[2]
	generateCompressedFile(outputFile, table, filename)
	defer f.Close()
}
