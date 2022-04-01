package main

import "io"

// Run brainfuck program, represented as a string
func executeProgram(s string, output io.Writer) {
	instructions := tokenize(s)
	array := initialArray(output)
	pointer := 0
	for pointer < len(instructions) {
		instructions[pointer].execute(&array, &pointer, &instructions)
	}
}
