package main

import (
	"flag"
	"fmt"
)

func catchUserInput() (string, byte, bool) {
	var debug bool

	filenamePtr := flag.String("file", "testInput.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of the puzzle do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		return *filenamePtr, 'a', debug
	case "b":
		return *filenamePtr, 'b', debug

	default:
		return *filenamePtr, 'z', debug
	}
}

func calcRange(calcChar rune, lowerLimit int, upperLimit int, debug bool) (int, int) {
	var newSize int
	var newLowerLimit, newUpperLimit int

	size := upperLimit - lowerLimit + 1
	if debug {
		fmt.Printf("Before calc... size: %d lowerLimit: %d upperLimit: %d\n", size, lowerLimit, upperLimit)
	}

	newSize = size / 2
	if calcChar == 'F' || calcChar == 'L' {
		newLowerLimit = lowerLimit
		newUpperLimit = lowerLimit + newSize - 1
	} else {
		newLowerLimit = lowerLimit + newSize
		newUpperLimit = upperLimit
	}

	if debug {
		fmt.Printf("After calc... size: %d lowerLimit: %d upperLimit: %d\n", newSize, newLowerLimit, newUpperLimit)
	}
	return newLowerLimit, newUpperLimit
}

func decodeSinglePass(boardingPass string, part byte, debug bool) int {
	if debug {
		fmt.Println("Processing: ", boardingPass)
	}

	//var rows int = 128
	var lowerLimit int = 0
	var upperLimit int = 127
	var rowNumber int = 0
	for _, rowChar := range boardingPass[0:7] {
		if debug {
			fmt.Printf("Row Instruction: %c\n", rowChar)
		}
		lowerLimit, upperLimit = calcRange(rowChar, lowerLimit, upperLimit, debug)
	}
	rowNumber = lowerLimit

	lowerLimit = 0
	upperLimit = 7
	var columnNumber int = 0
	for _, columnChar := range boardingPass[7:] {
		if debug {
			fmt.Printf("Column Instruction: %c\n", columnChar)
		}
		lowerLimit, upperLimit = calcRange(columnChar, lowerLimit, upperLimit, debug)
	}
	columnNumber = lowerLimit

	return (rowNumber * 8) + columnNumber
}

func processSinglePersonAnswers(answers string, answersStore map[byte]bool, part byte, debug bool) {
	// Not sure if this is the right thing to do
	// Build a map of the questions
	// Record the yes answers in the map
	// return when done

	for _, answer := range answers {
		answersStore[byte(answer)] = true
	}
}

func countAllAnswers(answersStore map[byte]int, numOfPeople int) int {
	var groupCount = 0

	for _, value := range answersStore {
		if value == numOfPeople {
			groupCount++
		}
	}
	return groupCount
}

func countAnyAnswers(answersStore map[byte]bool) int {
	var groupCount = 0

	for _, value := range answersStore {
		if value {
			groupCount++
		}
	}
	return groupCount
}

func processGroupAnswers(filename string, part byte, debug bool) int {
	var answersStore map[byte]bool
	var sumOfCounts int = 0

	answersStore = make(map[byte]bool)

	puzzleInput, _ := readFile(filename)

	for _, singlePersonAnswers := range puzzleInput {
		if debug {
			fmt.Println(singlePersonAnswers)
		}

		if singlePersonAnswers == "" {
			// Onto the next group
			// Sum up the total number of answers we have for this group
			if debug {
				fmt.Println(answersStore)
			}

			sumOfCounts += countAnyAnswers(answersStore)
			// Clear the answers store and start again
			answersStore = make(map[byte]bool)
		}

		processSinglePersonAnswers(singlePersonAnswers, answersStore, part, debug)
	}

	// count the last entry in case there wasn't a blank line before EOF
	sumOfCounts += countAnyAnswers(answersStore)

	return sumOfCounts
}

// Main routine
func main() {
	filenamePtr, execPart, debug := catchUserInput()
	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("Sum of answer counts:", processGroupAnswers(filenamePtr, execPart, debug))
	} else {
		fmt.Println("Not implemented yet")
	}

	if debug {
		var answersStore map[byte]bool
		answersStore = make(map[byte]bool)
		fmt.Printf("Answers: %s \n", "wzaopvscxknyjtiul")
		processSinglePersonAnswers("wzaopvscxknyjtiul", answersStore, execPart, debug)
		fmt.Println(answersStore)
	}

}