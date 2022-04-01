package main

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		expected []instruction
	}{
		{
			"First",
			"+-[]><.",
			[]instruction{startProgram{}, plus{}, minus{}, startCycle{}, endCycle{},
				increment{}, decrement{}, printValue{}, endProgram{}},
		},
		{
			"Second",
			"+++-.",
			[]instruction{startProgram{}, plus{}, plus{}, plus{}, minus{}, printValue{}, endProgram{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tokenize(tt.command)

			if len(result) != len(tt.expected) {
				t.Errorf("%s failed, expected length: %v, but found: %v",
					tt.name, len(tt.expected), len(result))
			} else {

				for i := range result {
					if reflect.TypeOf(result[i]) != reflect.TypeOf(tt.expected[i]) {
						t.Errorf("%s failed, expected: %s, but found: %s",
							tt.name, reflect.TypeOf(tt.expected[i]), reflect.TypeOf(result[i]))
					}
				}
			}
		})
	}
}
