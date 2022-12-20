package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getStackAndMoveInputs(reader io.Reader) ([]string, []string) {
	stackInputs := make([]string, 0)
	moveInputs := make([]string, 0)

	isStackInput := true

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			isStackInput = false
		} else if isStackInput {
			stackInputs = append(stackInputs, line)
		} else {
			moveInputs = append(moveInputs, line)
		}
	}

	return stackInputs, moveInputs
}

func convertStackInputToCargoShip(stackInputs []string) CargoShip {
	regularExpression := regexp.MustCompile("[0-9]+")
	numberOfStackInputLines := len(stackInputs)
	stackNumbersExtracted := regularExpression.FindAllString(stackInputs[numberOfStackInputLines-1], -1)
	numberOfStacks := len(stackNumbersExtracted)
	stackList := make([]*Stack, numberOfStacks)

	for index, _ := range stackList {
		stackList[index] = &Stack{}
	}

	for lineNumber := numberOfStackInputLines - 2; lineNumber >= 0; lineNumber-- {
		lineInput := stackInputs[lineNumber]
		for index, character := range lineInput {
			if index%4 == 1 && string(character) != " " {
				stackNumber := index / 4
				stack := stackList[stackNumber]
				stack.Push(string(character))
			}
		}
	}

	return CargoShip{
		StackList:      stackList,
		NumberOfStacks: numberOfStacks,
	}

}

func convertMoveInputsToRearrangementProcedure(moveInputs []string) RearrangementProcedure {
	regularExpression := regexp.MustCompile("[0-9]+")

	rearrangementProcedure := make(RearrangementProcedure, 0)

	for _, moveInput := range moveInputs {
		moveInputExtractedNumbers := regularExpression.FindAllString(moveInput, -1)

		numberOfCrates, _ := strconv.Atoi(moveInputExtractedNumbers[0])
		current, _ := strconv.Atoi(moveInputExtractedNumbers[1])
		next, _ := strconv.Atoi(moveInputExtractedNumbers[2])
		rearrangementProcedure = append(rearrangementProcedure, &Move{
			NumberOfCrates: numberOfCrates,
			Current:        current,
			Next:           next,
		})
	}

	return rearrangementProcedure
}

func main() {
	file, err := os.Open("/home/ec2-user/go/src/github.com/iamwillzhu/adventofcode2022day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stackInputs, moveInputs := getStackAndMoveInputs(file)

	cargoShipV1 := convertStackInputToCargoShip(stackInputs)
	rearrangementProcedure := convertMoveInputsToRearrangementProcedure(moveInputs)

	fmt.Println("Performing CrateMover 9000 moves...")
	for _, move := range rearrangementProcedure {
		cargoShipV1.PerformCrateMoverOperationV1(move)
	}

	topStackValuesPartOne := make([]string, 0)
	for _, stack := range cargoShipV1.StackList {
		if topStackValue, exist := stack.Top(); exist {
			topStackValuesPartOne = append(topStackValuesPartOne, topStackValue)
		}
	}

	resultPartOne := strings.Join(topStackValuesPartOne, "")

	fmt.Printf("The crates that end up at the top of each stack for part one is : %s\n", resultPartOne)

	cargoShipV2 := convertStackInputToCargoShip(stackInputs)

	fmt.Println("Performing CrateMover 9001 moves...")
	for _, move := range rearrangementProcedure {
		cargoShipV2.PerformCrateMoverOperationV2(move)
	}

	topStackValuesPartTwo := make([]string, 0)
	for _, stack := range cargoShipV2.StackList {
		if topStackValue, exist := stack.Top(); exist {
			topStackValuesPartTwo = append(topStackValuesPartTwo, topStackValue)
		}
	}

	resultPartTwo := strings.Join(topStackValuesPartTwo, "")

	fmt.Printf("The crates that end up at the top of each stack for part two is : %s\n", resultPartTwo)

}
