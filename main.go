package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("stack.txt")

	var s Stack

	var lines []string

	var currentValue string

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
			} else if splitcmd[0] == "Pop" {
				currentValue = s.Pop()
			} else if splitcmd[0] == "dupe" {
				s.Dupe()
			} else if splitcmd[0] == "PushCurrent" {
				s.Push(currentValue)
			} else if splitcmd[0] == "Add" {
				firstNumber, error1 := strconv.ParseInt(s.Pop(), 10, 64)
				secondNumber, error2 := strconv.ParseInt(s.Pop(), 10, 64)
				if error1 != nil || error2 != nil {
					log.Fatal("Invalid Addition")
				} else {
					s.Push(strconv.FormatInt(secondNumber+firstNumber, 10))
				}

			} else if splitcmd[0] == "Subtract" {
				firstNumber, error1 := strconv.ParseInt(s.Pop(), 10, 64)
				secondNumber, error2 := strconv.ParseInt(s.Pop(), 10, 64)
				if error1 != nil || error2 != nil {
					log.Fatal("Invalid Subtraction")
				} else {
					s.Push(strconv.FormatInt(secondNumber-firstNumber, 10))
				}
			} else if splitcmd[0] == "Multiply" {
				firstNumber, error1 := strconv.ParseInt(s.Pop(), 10, 64)
				secondNumber, error2 := strconv.ParseInt(s.Pop(), 10, 64)
				if error1 != nil || error2 != nil {
					log.Fatal("Invalid Multiplication")
				} else {
					s.Push(strconv.FormatInt(secondNumber*firstNumber, 10))
				}
			} else if splitcmd[0] == "Divide" {
				firstNumber, error1 := strconv.ParseInt(s.Pop(), 10, 64)
				secondNumber, error2 := strconv.ParseInt(s.Pop(), 10, 64)
				if error1 != nil || error2 != nil {
					log.Fatal("Invalid Division")
				} else {
					s.Push(strconv.FormatInt(secondNumber/firstNumber, 10))
				}
			} else if splitcmd[0] == "Exponentiate" {
				firstNumber, error1 := strconv.ParseInt(s.Pop(), 10, 64)
				secondNumber, error2 := strconv.ParseInt(s.Pop(), 10, 64)
				if error1 != nil || error2 != nil {
					log.Fatal("Invalid Exponentiation")
				} else {
					s.Push(strconv.FormatInt(int64(math.Pow(float64(secondNumber), float64(firstNumber))), 10))
				}
			} else if splitcmd[0] == "stdin" {
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					s.Push(scanner.Text())
				}
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
		} else if command == "PRINT" {
			return "Print"
		} else if command == "^" {
			return "Exponentiate"
		} else if command == "STDIN" {
			return "stdin"
		} else if command == "POP" {
			return "Pop"
		} else if command == "PUSH" {
			return "PushCurrent"
		} else if command == "DUPE" {
			return "dupe"
		} else {
			return "Push>" + command
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

func (s *Stack) Dupe() {
	cmd := s.Pop()
	s.Push(cmd)
	s.Push(cmd)
}
