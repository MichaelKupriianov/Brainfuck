package main

import (
	"bytes"
	"testing"
)

func TestPlus(t *testing.T) {
	tests := []struct {
		name     string
		a        array
		expected int
	}{
		{
			"First",
			array{cells: []int{1, 2, 3}, pointer: 1},
			3,
		},
		{
			"Second",
			array{cells: []int{1, 0, 2}, pointer: 0},
			2,
		},
		{
			"Third",
			array{cells: []int{1, 0}, pointer: 1},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pointer := 1
			i := plus{}

			i.execute(&tt.a, &pointer, nil)

			if tt.a.cells[tt.a.pointer] != tt.expected {
				t.Errorf("%s failed, expected: %v in %v cell, but found: %v",
					tt.name, tt.expected, tt.a.pointer, tt.a.cells[tt.a.pointer])
			}
		})
	}
}

func TestMinus(t *testing.T) {
	tests := []struct {
		name     string
		a        array
		expected int
	}{
		{
			"First",
			array{cells: []int{1, 2, 3}, pointer: 1},
			1,
		},
		{
			"Second",
			array{cells: []int{1, 0, 2}, pointer: 0},
			0,
		},
		{
			"Third",
			array{cells: []int{1, 0}, pointer: 1},
			-1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pointer := 1
			i := minus{}

			i.execute(&tt.a, &pointer, nil)

			if tt.a.cells[tt.a.pointer] != tt.expected {
				t.Errorf("%s failed, expected: %v in %v cell, but found: %v",
					tt.name, tt.expected, tt.a.pointer, tt.a.cells[tt.a.pointer])
			}
		})
	}
}

func TestPrint(t *testing.T) {
	output := bytes.NewBuffer([]byte{})

	tests := []struct {
		name     string
		a        array
		expected string
	}{
		{
			"First",
			array{[]int{1, 2, 3}, 1, output},
			"2 ",
		},
		{
			"Second",
			array{[]int{1, 0, 2}, 0, output},
			"1 ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pointer := 1
			i := printValue{}

			i.execute(&tt.a, &pointer, nil)

			if tt.expected != output.String() {
				t.Errorf("%s failed, expected: %s, but found: %s",
					tt.name, tt.expected, output.String())
			}

			output.Reset()
		})
	}
}

func TestCycle(t *testing.T) {
	tests := []struct {
		name     string
		a        array
		pointer  int
		program  []instruction
		expected int
	}{
		{
			"First",
			array{cells: []int{0, 2, 0}, pointer: 1},
			0,
			[]instruction{startCycle{}, increment{}, plus{}, decrement{}, minus{}, endCycle{}, endProgram{}},
			2,
		},
		{
			"Second",
			array{cells: []int{2, 2, 0}, pointer: 0},
			0,
			[]instruction{
				startCycle{}, increment{}, plus{},
				startCycle{}, increment{}, plus{}, decrement{}, minus{}, endCycle{},
				decrement{}, minus{}, endCycle{}, endProgram{},
			},
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for tt.pointer < len(tt.program) {
				tt.program[tt.pointer].execute(&tt.a, &tt.pointer, &tt.program)
			}

			if tt.expected != tt.a.cells[2] {
				t.Errorf("%s failed, expected: %v, but found: %v",
					tt.name, tt.expected, tt.a.cells[2])
			}
		})
	}
}

func TestIncrement(t *testing.T) {
	tests := []struct {
		name     string
		a        array
		expected int
	}{
		{
			"First",
			array{cells: []int{1, 2, 3}, pointer: 1},
			2,
		},
		{
			"Second",
			array{cells: []int{1, 0, 2}, pointer: 0},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pointer := 1
			i := increment{}

			i.execute(&tt.a, &pointer, nil)

			if tt.a.pointer != tt.expected {
				t.Errorf("%s failed, expected: %v, but found: %v",
					tt.name, tt.expected, tt.a.pointer)
			}
		})
	}
}

func TestDecrement(t *testing.T) {
	tests := []struct {
		name     string
		a        array
		expected int
	}{
		{
			"First",
			array{cells: []int{1, 2, 3}, pointer: 1},
			0,
		},
		{
			"Second",
			array{cells: []int{1, 0, 2}, pointer: 0},
			-1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pointer := 1
			i := decrement{}

			i.execute(&tt.a, &pointer, nil)

			if tt.a.pointer != tt.expected {
				t.Errorf("%s failed, expected: %v, but found: %v",
					tt.name, tt.expected, tt.a.pointer)
			}
		})
	}
}
