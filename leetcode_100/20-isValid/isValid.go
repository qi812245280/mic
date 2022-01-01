package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	base := map[string]string{"(": "(", "{": "{", "[": "[", ")": "(", "}": "{", "]": "["}
	idx := 0
	stack := make([]string, 0)
	for _, c := range s {
		str := string(c)
		if strings.Contains("({[", str) {
			stack = append(stack, str)
			idx++
		}
		if strings.Contains(")}]", str) {
			if idx == 0 {
				return false
			}
			if stack[idx-1] != base[str] {
				return false
			}
			stack = stack[:len(stack)-1]
			idx--
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
