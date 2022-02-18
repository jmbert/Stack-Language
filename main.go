package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("stack.txt")

	var s Stack

	var lines []string

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		var cmd = strings.Split(lines[i], " ")

		for i := 0; i < len(cmd); i++ {

			var stackCommands = GenerateCommands(cmd[i])

			splitcmd := strings.Split(stackCommands, ">")
			if splitcmd[0] == "Push" {
				s.Push(splitcmd[1])
			} else if splitcmd[0] == "Add" {
				firstNumber, _ := strconv.ParseInt(s.Pop(), 10, 64)
				secondNumber, _ := strconv.ParseInt(s.Pop(), 10, 64)
				s.Push(strconv.FormatInt(secondNumber+firstNumber, 10))
			} else if splitcmd[0] == "Subtract" {
				firstNumber, _ := strconv.ParseInt(s.Pop(), 10, 64)
				secondNumber, _ := strconv.ParseInt(s.Pop(), 10, 64)
				s.Push(strconv.FormatInt(secondNumber-firstNumber, 10))
			} else if splitcmd[0] == "Multiply" {
				firstNumber, _ := strconv.ParseInt(s.Pop(), 10, 64)
				secondNumber, _ := strconv.ParseInt(s.Pop(), 10, 64)
				s.Push(strconv.FormatInt(secondNumber*firstNumber, 10))
			} else if splitcmd[0] == "Divide" {
				firstNumber, _ := strconv.ParseInt(s.Pop(), 10, 64)
				secondNumber, _ := strconv.ParseInt(s.Pop(), 10, 64)
				s.Push(strconv.FormatInt(secondNumber/firstNumber, 10))
			} else if splitcmd[0] == "Exponentiate" {
				firstNumber, _ := strconv.ParseInt(s.Pop(), 10, 64)
				secondNumber, _ := strconv.ParseInt(s.Pop(), 10, 64)
				s.Push(strconv.FormatInt(int64(math.Pow(float64(secondNumber), float64(firstNumber))), 10))
			} else if splitcmd[0] == "Print" {
				topCmd := s.Pop()
				println(topCmd)
				s.Push(topCmd)
			}
		}
	}
}

func GenerateCommands(command string) string {
	_, error := strconv.ParseInt(command, 10, 64)

	if error != nil {
		if command == "+" {
			return "Add"
		} else if command == "-" {
			return "Subtract"
		} else if command == "*" {
			return "Multiply"
		} else if command == "/" {
			return "Divide"
		} else if command == "print" {
			return "Print"
		} else if command == "^" {
			return "Exponentiate"
		} else {
			return ""
		}
	} else {
		return "Push>" + command
	}
}

type Stack struct {
	stack []string
}

func (s *Stack) Push(command string) {
	s.stack = append(s.stack, command)
}

func (s *Stack) Pop() string {
	var poppedStack []string
	var lastElement = s.stack[len(s.stack)-1]

	for i := 0; i < len(s.stack)-1; i++ {
		poppedStack = append(poppedStack, s.stack[i])
	}
	s.stack = poppedStack

	return lastElement
}
