package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		if str != nil {
			fmt.Println(str)
		}
	}()
	code := read()
	executeProgram(code)
}
