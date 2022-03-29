package main

type array struct {
	cells   [40000]int
	pointer int
}

func initialArray() array {
	return array{pointer: 20000}
}

func (a array) getValue() int {
	return a.cells[a.pointer]
}
