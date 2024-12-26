package piscine

import (
	"fmt"
	"os"
)

const (
	MinInt64 = -1 << 63
	MaxInt64 = 1<<63 - 1
)

func main() {
	input := os.Args[1:]
	if len(input) != 3 {
		return
	}

	op := input[1]

	n1, ok := Atoi(input[0])
	if !ok {
		return
	}

	n2, ok := Atoi(input[2])
	if !ok {
		return
	}

	output := do(op)(n1, n2)
	if output == "" {
		return
	}

	Println(output)
}

func do(op string) func(a, b int) string {
	switch op {
	case "+":
		return add
	case "-":
		return substract
	case "*":
		return multiply
	case "/":
		return divide
	case "%":
		return modulo
	default:
		return nothing
	}
}

func nothing(a, b int) string {
	return ""
}

func add(a, b int) string {
	if a > 0 && b > MaxInt64-a { // 5  4 > 9-5
		return ""
	}
	if a < 0 && b < MinInt64-a { // -5 -5 < -10 - -5
		return ""
	}
	return Itoa(a + b)
}

func substract(a, b int) string {
	if a > 0 && b < a-MaxInt64 { // 5 - -4 < 5-9
		return ""
	}
	if a < 0 && b > a-MinInt64 { // -5 - 5 > 5
		return ""
	}
	return Itoa(a - b)
}

func multiply(a, b int) string {
	if a == 0 || b == 0 {
		return "0"
	}
	c := a * b
	if c/a != b {
		return ""
	}
	if c/b != a {
		return ""
	}
	return Itoa(c)
}

func divide(a, b int) string {
	if b == 0 {
		return "No division by 0"
	}
	return Itoa(a / b)
}

func modulo(a, b int) string {
	if b == 0 {
		return "No modulo by 0"
	}
	return Itoa(a % b)
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func Atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	if s == "0" {
		return 0, true
	}

	negative := s[0] == '-'
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

	var n int

	for _, r := range s {
		if !isDigit(r) {
			return 0, false
		}
		n *= 10
		n -= int(r - '0')
		fmt.Println(n, "??")
	}

	if negative {
		if n < 0 {
			return n, true
		}
		return 0, false
	}

	n = -n

	if n > 0 {
		return n, true
	}

	return 0, false
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	negative := n < 0

	var runes []rune

	for n != 0 {
		mod := n % 10
		if mod < 0 {
			mod = -mod
		}
		runes = append([]rune{rune(mod) + '0'}, runes...)
		n /= 10
	}

	if negative {
		runes = append([]rune{'-'}, runes...)
	}

	return string(runes)
}

func Println(s string) {
	os.Stdout.WriteString(s)
	os.Stdout.WriteString("\n")
}
