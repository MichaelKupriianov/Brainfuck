package main

func tokenize(s string) []instruction {
	var result []instruction
	result = append(result, startProgram{})
	for i := range s {
		switch s[i] {
		case '+':
			result = append(result, plus{})
		case '-':
			result = append(result, minus{})
		case '.':
			result = append(result, print{})
		case '[':
			result = append(result, startCycle{})
		case ']':
			result = append(result, endCycle{})
		case '>':
			result = append(result, increment{})
		case '<':
			result = append(result, decrement{})
		}
	}
	result = append(result, endProgram{})
	return result
}
