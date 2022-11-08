package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var FA *FiniteAutomaton
var commandDict = make(map[int]func())
var cin = bufio.NewReader(os.Stdin)
var sequence string

func printStates() {

	fmt.Print("All FA states: ")

	for _, st := range FA.States {

		fmt.Print(st, " ")
	}

	fmt.Println()
}

func printAlphabet() {

	fmt.Print("The FA alphabet: ")

	for _, term := range FA.Alphabet {

		fmt.Print(term, " ")
	}

	fmt.Println()
}

func printTransitions() {

	fmt.Print("All FA transitions:\n")

	for _, tr := range FA.Transitions {

		fmt.Print(tr.String(), "\n")
	}

	fmt.Println()
}

func printInitialState() {

	fmt.Println("The FA initial state: ", FA.InitialState)
}

func printFinalStates() {

	fmt.Print("The FA final states: ")

	for _, st := range FA.FinalStates {

		fmt.Print(st, " ")
	}

	fmt.Println()
}

func checkSequence() {

	fmt.Print("\nEnter sequence: ")

	readBytes, _ := cin.ReadBytes('\n')

	sequence = string(readBytes[:len(readBytes)-2])

	if FA.Accepts(sequence) {
		fmt.Println("Sequence accepted!")

	} else {
		fmt.Println("Sequence not accepted!")
	}
}

func printMenu() {

	fmt.Println("1. Display the set of states")
	fmt.Println("2. Display the alphabet")
	fmt.Println("3. Display all the transitions")
	fmt.Println("4. Display the initial state")
	fmt.Println("5. Display the final states")
	fmt.Println("6. Check if a sequence is accepted by the FA")
	fmt.Println("0. Exit")
}

func main() {

	var err error
	var commandNo int

	FA = new(FiniteAutomaton)

	err = FA.ReadFA("Lab4/FA.in")

	if err != nil {
		log.Fatalln(err)
		return
	}

	commandDict[1] = printStates
	commandDict[2] = printAlphabet
	commandDict[3] = printTransitions
	commandDict[4] = printInitialState
	commandDict[5] = printFinalStates
	commandDict[6] = checkSequence

	printMenu()

	exit := false

	for !exit {

		fmt.Print("\nFA>> ")

		readString, _ := cin.ReadBytes('\n')

		readString = readString[:len(readString)-2]

		commandNo, err = strconv.Atoi(string(readString))

		if err != nil {
			log.Println("invalid command")
			continue
		}

		if commandNo < 0 || commandNo > 6 {
			log.Println("invalid command")
			continue
		}

		if commandNo == 0 {
			exit = true

		} else {
			commandDict[commandNo]()
		}
	}
}
