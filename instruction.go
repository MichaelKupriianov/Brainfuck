package main

type instruction interface {
	execute(*array, *int, *[]instruction)
}

// Increases the value in current cell by 1, corresponds to the "+" symbol in brainfuck syntax
type plus struct{}

func (plus) execute(a *array, p *int, code *[]instruction) {
	a.cells[a.pointer]++
	*p++
}

// Decreases the value in current cell by 1, corresponds to the "-" symbol in brainfuck syntax
type minus struct{}

func (minus) execute(a *array, p *int, code *[]instruction) {
	a.cells[a.pointer]--
	*p++
}

// Print the value in current cell, corresponds to the "." symbol in brainfuck syntax
type printValue struct{}

func (printValue) execute(a *array, p *int, code *[]instruction) {
	a.printValue()
	*p++
}

// Terminate the cycle if value in current cell is 0, corresponds to the "[" symbol in brainfuck syntax
type startCycle struct{}

func (startCycle) execute(a *array, p *int, code *[]instruction) {
	if a.getValue() == 0 {
		counter := 0
		for {
			*p++
			switch (*code)[*p].(type) {
			case endProgram:
				panic("Incorrect program. There are more opening brackets than closed.")
			case startCycle:
				counter++
			case endCycle:
				if counter == 0 {
					*p++
					return
				}
				counter--
			}
		}
	} else {
		*p++
	}
}

// Continue the cycle if value in current cell is not 0, corresponds to the "]" symbol in brainfuck syntax
type endCycle struct{}

func (endCycle) execute(a *array, p *int, code *[]instruction) {
	if a.getValue() != 0 {
		counter := 0
		for {
			*p--
			switch (*code)[*p].(type) {
			case startProgram:
				panic("Incorrect program. There are more closing brackets than opening.")
			case endCycle:
				counter++
			case startCycle:
				if counter == 0 {
					*p++
					return
				}
				counter--
			}
		}
	} else {
		*p++
	}
}

// Move to next cell, corresponds to the ">" symbol in brainfuck syntax
type increment struct{}

func (increment) execute(a *array, p *int, code *[]instruction) {
	a.pointer++
	*p++
}

// Move to previous cell, corresponds to the "<" symbol in brainfuck syntax
type decrement struct{}

func (decrement) execute(a *array, p *int, code *[]instruction) {
	a.pointer--
	*p++
}

// Marks start of the program
type startProgram struct{}

func (startProgram) execute(a *array, p *int, code *[]instruction) {
	*p++
}

// Marks end of the program
type endProgram struct{}

func (endProgram) execute(a *array, p *int, code *[]instruction) {
	*p++
}
