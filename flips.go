package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Opens the defined input file and iterates over each stack of pancakes to determine
// the number of flips required
func main() {

	filePath := flag.String("filePath", "/tmp/pancake_data", "The path to the input file containing the pancake stacks")
	flag.Parse()

	// Attempt to open to the input file and ensure it was done successfully
	file, err := os.Open(*filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Increment the location of the scanner to get to the stacks
	scanner.Scan()

	caseNum := 1
	// For each stack of pancakes, find the optimal number of flips and print the
	// result in the required format
	for scanner.Scan() {
		fmt.Printf("Case #%d: %d\n", caseNum, countFlips(scanner.Text()))
		caseNum++;
	}

}

// countFlips takes a string representing a stack of pancakes and determines the
// optimal number of flips to perform to get the entire stack "+" side up.
// It returns the the number of flips performed on the stack.
func countFlips(pancakeStack string) int {

	flipCount := 0
	currentStackSign := ""

	// Split the input string on each character
	stackSlice := strings.Split(pancakeStack, "")

	for i := 0 ; i < len(stackSlice); i++ {

		currentPancakeSign := stackSlice[i];

		// Validate inputs
		if (!strings.Contains(currentPancakeSign, "-") && !strings.Contains(currentPancakeSign, "+")){
			panic("Invlaid inputs in the pancake_data file")
		}

		if(strings.Compare(currentStackSign, "") == 0 ){
			currentStackSign = currentPancakeSign

		// Flip the stack and increment our counter when
		// the pancake we are inspecting isn't the same as the previous
		// pancakes on the stack.
		} else if currentPancakeSign != currentStackSign {
			currentStackSign = currentPancakeSign
			flipCount = flipCount + 1
		}
	}

	// If we've normalized the stack to all be "-" do one final flip
	// to get our result
	if(strings.Compare(currentStackSign, "-") == 0 ){
		flipCount = flipCount + 1
	}

	return flipCount
}

// If we encounter an unexpected error, terminate the program
func check(e error) {
	if e != nil {
		panic(e)
	}
}