package main

import (
	"bytes"
	"io"
	"testing"
)

func TestExecuteProgram(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		expected string
		output   io.Writer
	}{
		{
			"Hello world",
			"++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.",
			"72 101 108 108 111 32 87 111 114 108 100 33 10 ",
			bytes.NewBuffer([]byte{}),
		},
		{
			"Simple loop",
			"[].",
			"0 ",
			bytes.NewBuffer([]byte{}),
		},
		{
			"Difficult loop",
			"+++++[.-]",
			"5 4 3 2 1 ",
			bytes.NewBuffer([]byte{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executeProgram(tt.command, tt.output)

			if tt.expected != tt.output.(*bytes.Buffer).String() {
				t.Errorf("%s failed, expected: %s, but found: %s",
					tt.name, tt.expected, tt.output.(*bytes.Buffer).String())
			}
		})
	}
}
