package ts

import (
	"fmt"
	"os"
)

// Используем стэк
func matchBrackets(exp string) bool {
	runes := []rune(exp)
	var opened []rune
	ptr := -1
	for _, c := range runes {
		if c == '(' || c == '[' || c == '{' {
			opened = append(opened, c)
			ptr++
		} else if c == ')' {
			if ptr < 0 || opened[ptr] != '(' {
				return false
			}
			opened = opened[:len(opened)-1]
			ptr--
		} else if c == ']' {
			if ptr < 0 || opened[ptr] != '[' {
				return false
			}
			opened = opened[:len(opened)-1]
			ptr--
		} else if c == '}' {
			if ptr < 0 || opened[ptr] != '{' {
				return false
			}
			opened = opened[:len(opened)-1]
			ptr--
		}
	}
	return len(opened) == 0
}

func Brackets() {
	if len(os.Args) == 1 {
		fmt.Println()
	} else {
		for _, v := range os.Args[1:] {
			if matchBrackets(v) {
				fmt.Println("OK")
			} else {
				fmt.Println("Error")
			}
		}
	}
}
