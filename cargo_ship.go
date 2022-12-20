package main

type CargoShip struct {
	StackList      []*Stack
	NumberOfStacks int
}

type Move struct {
	NumberOfCrates int
	Current        int
	Next           int
}

type RearrangementProcedure []*Move

func (c *CargoShip) PerformCrateMoverOperationV1(move *Move) {
	currentStack := c.StackList[move.Current-1]
	nextStack := c.StackList[move.Next-1]

	for crateNumber := 1; crateNumber <= move.NumberOfCrates; crateNumber++ {
		crateToMove, _ := currentStack.Pop()
		nextStack.Push(crateToMove)
	}
}

func (c *CargoShip) PerformCrateMoverOperationV2(move *Move) {
	currentStack := c.StackList[move.Current-1]
	nextStack := c.StackList[move.Next-1]

	temporaryStack := &Stack{}
	for crateNumber := 1; crateNumber <= move.NumberOfCrates; crateNumber++ {
		crateToMove, _ := currentStack.Pop()
		temporaryStack.Push(crateToMove)
	}

	for !temporaryStack.IsEmpty() {
		crateToMove, _ := temporaryStack.Pop()
		nextStack.Push(crateToMove)
	}
}
