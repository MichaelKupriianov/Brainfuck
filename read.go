package main

import "fmt"

func read() string {
	var program string
	fmt.Scan(&program)

	for i := range program {
		if program[i] != '+' && program[i] != '-' && program[i] != '.' &&
			program[i] != '[' && program[i] != ']' && program[i] != '>' && program[i] != '<' {
			panic("Found a symbol that has no meaning in BrainFuck")
		}
	}
	return program
}
