package ts

import (
	"os"

	"github.com/01-edu/z01"
)

func BrainFuck() {
	if len(os.Args) != 2 {
		return
	}

	program := []byte(os.Args[1])
	var cursor int

	var memory [2048]byte
	var pointer int

	for cursor >= 0 && cursor < len(program) {
		switch program[cursor] {
		case '>':
			// Increment the pointer
			pointer++
		case '<':
			// decrement the pointes
			pointer--
		case '+':
			// increment the pointed byte
			memory[pointer]++
		case '-':
			// decrement the pointed byte
			memory[pointer]--
		case '.':
			// print the pointed byte on std output
			z01.PrintRune(rune(memory[pointer]))

		case '[':
			// go to the matching ']' if the pointed byte is 0 (while start)
			stackCount := 0
			if memory[pointer] == 0 {
				for cursor < len(program) && (program[cursor] != byte(']') || stackCount > 1) {
					if program[cursor] == byte('[') {
						stackCount++
					} else if program[cursor] == byte(']') {
						stackCount--
					}
					cursor++
				}
			}
		case ']':
			// go to the matching '[' if the pointed byte is not 0 (while end)
			stackCount := 0
			if memory[pointer] != 0 {
				for cursor >= 0 && (program[cursor] != byte('[') || stackCount > 1) {
					if program[cursor] == byte(']') {
						stackCount++
					} else if program[cursor] == byte('[') {
						stackCount--
					}
					cursor--
				}
			}
		}
		cursor++
	}
}
