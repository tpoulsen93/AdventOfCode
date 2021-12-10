package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// I pulled the stack implementation from https://www.educative.io/edpresso/how-to-implement-a-stack-in-golang
type Stack []rune

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(char rune) {
	*s = append(*s, char) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// check if a rune is an opening bracket
func isOpener(target rune) bool {
	return strings.Contains("({[<", string(target))
}

// check if the popped opening bracket matches the target: <>{}[]()
func isMatch(target rune, popped rune) bool {
	switch popped {
	case '(':
		if target == ')' {
			return true
		} else {
			return false
		}
	case '[':
		if target == ']' {
			return true
		} else {
			return false
		}
	case '{':
		if target == '}' {
			return true
		} else {
			return false
		}
	case '<':
		if target == '>' {
			return true
		} else {
			return false
		}
	default:
		return false // this line should never be reached...
	}
}

func main() {
	// get the file and pass it to the scanner
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// build the structure for the values of characters
	bounties := make(map[rune]int)
	bounties[')'] = 3
	bounties[']'] = 57
	bounties['}'] = 1197
	bounties['>'] = 25137

	// keep score
	score := 0

	// use a stack for the brackets
	var stack Stack

	// loop through each line going through the input
	for scanner.Scan() {
		line := scanner.Text()

		// loop through the line adding brackets until a corrupted
		// bracket is found or we get to the end of the line
		for _, curr := range line {
			if isOpener(curr) {
				stack.Push(curr)
			} else {
				// the current rune must be a closing bracket
				popped, state := stack.Pop()

				if state && !isMatch(curr, popped) {
					score += bounties[curr]
					break
				}
			}
		}
		// empty the stack before the next iteration
		stack = stack[:0]
	}
	fmt.Println("Score:", score)
}
