package main

import (
	"fmt"
	"io"
)

type array struct {
	cells   []int
	pointer int
	output  io.Writer
}

func initialArray(output io.Writer) array {
	return array{make([]int, 40000), 20000, output}
}

func (a array) getValue() int {
	return a.cells[a.pointer]
}

func (a array) printValue() {
	fmt.Fprint(a.output, a.cells[a.pointer], " ")
}
