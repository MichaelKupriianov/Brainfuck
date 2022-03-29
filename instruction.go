package main

import "fmt"

type instruction interface {
	execute(*array, *int, *[]instruction)
}

type plus struct{}

func (plus) execute(a *array, p *int, code *[]instruction) {
	a.cells[a.pointer]++
	*p++
}

type minus struct{}

func (minus) execute(a *array, p *int, code *[]instruction) {
	a.cells[a.pointer]--
	*p++
}

type print struct{}

func (print) execute(a *array, p *int, code *[]instruction) {
	fmt.Print(a.cells[a.pointer], " ")
	*p++
}

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

type increment struct{}

func (increment) execute(a *array, p *int, code *[]instruction) {
	a.pointer++
	*p++
}

type decrement struct{}

func (decrement) execute(a *array, p *int, code *[]instruction) {
	a.pointer--
	*p++
}

type startProgram struct{}

func (startProgram) execute(a *array, p *int, code *[]instruction) {
	*p++
}

type endProgram struct{}

func (endProgram) execute(a *array, p *int, code *[]instruction) {
	*p++
}
