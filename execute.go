package main

// Run brainfuck program, represented as a string
func executeProgram(s string) {
	instructions := tokenize(s)
	array := initialArray()
	pointer := 0
	for pointer < len(instructions) {
		instructions[pointer].execute(&array, &pointer, &instructions)
	}
}
