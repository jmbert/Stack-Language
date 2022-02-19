package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	//fScanner := bufio.NewScanner(os.Stdin)

	//fScanner.Scan()
	//filename := fScanner.Text()

	f, err := os.Open("stack.txt")

	if err != nil {
		//log.Fatalf("No file %s exists. Try adding a complete path", filename)
	}

	var s Stack

	var lines []string

	var currentValue string

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		var cmd = strings.Split(lines[i], "\n")

		for j := 0; j < len(cmd); j++ {
			var stackCommands = GenerateCommands(cmd[j], s)

			splitcmd := strings.Split(stackCommands, "|")
			if splitcmd[0] == "Push" {
				s.Push(splitcmd[1])
			} else if splitcmd[0] == "Pop" {
				currentValue = s.Pop()
			} else if splitcmd[0] == "dupe" {
				s.Dupe()
			} else if splitcmd[0] == "PushCurrent" {
				s.Push(currentValue)
			} else if splitcmd[0] == "Add" {
				rawFirstNumber := s.Pop()
				rawSecondNumber := s.Pop()
				firstNumber, error1 := strconv.ParseInt(rawFirstNumber, 10, 64)
				secondNumber, error2 := strconv.ParseInt(rawSecondNumber, 10, 64)
				s.Push(rawSecondNumber)
				s.Push(rawFirstNumber)
				if error1 != nil || error2 != nil {
					log.Fatal("Invalid Addition")
				} else {
					s.Push(strconv.FormatInt(secondNumber+firstNumber, 10))
				}

			} else if splitcmd[0] == "Subtract" {
				rawFirstNumber := s.Pop()
				rawSecondNumber := s.Pop()
				firstNumber, error1 := strconv.ParseInt(rawFirstNumber, 10, 64)
				secondNumber, error2 := strconv.ParseInt(rawSecondNumber, 10, 64)
				s.Push(rawSecondNumber)
				s.Push(rawFirstNumber)
				if error1 != nil || error2 != nil {
					log.Fatal("Invalid Subtraction")
				} else {
					s.Push(strconv.FormatInt(secondNumber-firstNumber, 10))
				}
			} else if splitcmd[0] == "Multiply" {
				rawFirstNumber := s.Pop()
				rawSecondNumber := s.Pop()
				firstNumber, error1 := strconv.ParseInt(rawFirstNumber, 10, 64)
				secondNumber, error2 := strconv.ParseInt(rawSecondNumber, 10, 64)
				s.Push(rawSecondNumber)
				s.Push(rawFirstNumber)
				if error1 != nil || error2 != nil {
					log.Fatal("Invalid Multiplication")
				} else {
					s.Push(strconv.FormatInt(secondNumber*firstNumber, 10))
				}
			} else if splitcmd[0] == "Divide" {
				rawFirstNumber := s.Pop()
				rawSecondNumber := s.Pop()
				firstNumber, error1 := strconv.ParseInt(rawFirstNumber, 10, 64)
				secondNumber, error2 := strconv.ParseInt(rawSecondNumber, 10, 64)
				s.Push(rawSecondNumber)
				s.Push(rawFirstNumber)
				if error1 != nil || error2 != nil {
					log.Fatal("Invalid Division")
				} else {
					s.Push(strconv.FormatInt(secondNumber/firstNumber, 10))
				}
			} else if splitcmd[0] == "Exponentiate" {
				rawFirstNumber := s.Pop()
				rawSecondNumber := s.Pop()
				firstNumber, error1 := strconv.ParseInt(rawFirstNumber, 10, 64)
				secondNumber, error2 := strconv.ParseInt(rawSecondNumber, 10, 64)
				s.Push(rawSecondNumber)
				s.Push(rawFirstNumber)
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
			} else if splitcmd[0] == "input" {
				seperated := strings.Split(cmd[i], "(")

				//print(seperated[0])

				if len(seperated) != 2 {
					log.Fatal("Invalid input, add parentheses")
				} else {
					var input string
					//fmt.Print(seperated[1])
					fmt.Print(strings.Replace(seperated[1], ")", "", 1))
					fmt.Scanln(&input)
					s.Push(input)
				}

			} else if splitcmd[0] == "Print" {
				topCmd := s.Pop()
				println(topCmd)
				s.Push(topCmd)
			} else if splitcmd[0] == "PrintStr" {
				topCmd := s.Pop()
				var seperated = strings.Split(topCmd, ":")
				println(seperated[1])
				s.Push(topCmd)
			} else if splitcmd[0] == "clear" {
				for i := 0; i < len(s.stack); i++ {
					s.Pop()
				}
			} else if splitcmd[0] == "jump" {
				if len(splitcmd) != 2 {
					log.Fatal("Invalid jump, add parentheses")
				} else {
					usableCmd := splitcmd[1]

					if usableCmd == "__TOPREPLACE__" {
						usableCmd = s.Pop()
						s.Push(usableCmd)
					} else if usableCmd == "__TOP__" {
						usableCmd = s.Pop()
					}
					tryInt, err := strconv.ParseInt(usableCmd, 10, 64)

					if err != nil {
						log.Fatal("invalid input to jump")
					} else {
						i = int(tryInt - 1)
					}
				}
			} else if splitcmd[0] == "sleep" {
				if len(splitcmd) != 2 {
					log.Fatal("Invalid sleep, add parentheses")
				} else {
					usableCmd := splitcmd[1]
					tryInt, err := strconv.ParseInt(usableCmd, 10, 64)

					if err != nil {
						log.Fatal("invalid input to sleep")
					} else {
						time.Sleep(time.Duration(tryInt) * time.Second)
					}
				}
			} else if splitcmd[0] == "condjump" {
				condition := strings.Split(splitcmd[1], ":")

				topValStr := s.Pop()
				s.Push(topValStr)

				topVal, err := strconv.ParseInt(topValStr, 10, 64)

				if err != nil {
					log.Fatal("invalid top value")
				}

				if len(condition) != 2 {
					log.Fatal("invalid condition")
				} else {
					if TestCondition(condition[0], int(topVal)) {
						if len(splitcmd) != 2 {
							log.Fatal("Invalid jump, add parentheses")
						} else {
							usableCmd := condition[1]

							if usableCmd == "__TOPREPLACE__" {
								usableCmd = s.Pop()
								s.Push(usableCmd)
							} else if usableCmd == "__TOP__" {
								usableCmd = s.Pop()
							}
							tryInt, err := strconv.ParseInt(usableCmd, 10, 64)

							if err != nil {
								log.Fatal("invalid input to jump")
							} else {
								i = int(tryInt - 1)
							}
						}
					}
				}

			} else if splitcmd[0] == "length" {
				s.Push(strconv.FormatInt(int64(len(s.stack)), 10))
			} else if splitcmd[0] == "concat" {
				firstString := s.Pop()
				secondString := s.Pop()

				properString1, isString1 := CheckString(firstString)
				properString2, isString2 := CheckString(secondString)

				properString1 = strings.Replace(properString1, "\"", "", -1)
				properString2 = strings.Replace(properString2, "\"", "", -1)

				if !isString1 || !isString2 {
					log.Fatal("Invalid Concat")
				} else {

					s.Push("string:" + "\"" + properString2 + properString1 + "\"")
				}
			}
		}
	}
}

func GenerateCommands(command string, s Stack) string {
	seperated := strings.Split(command, "(")

	if command[0] != '"' {
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
				var top = s.Pop()
				s.Push(top)
				var seperated = strings.Split(top, ":")
				if len(seperated) != 1 {
					return "PrintStr|" + strings.Replace(seperated[1], ")", "", 1)
				} else {
					return "Print"
				}
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
			} else if seperated[0] == "INPUT" {
				return "input"
			} else if command == "CONCAT" {
				return "concat"
			} else if command == "CLEAR" {
				return "clear"
			} else if seperated[0] == "JUMP" {
				return "jump|" + strings.Replace(seperated[1], ")", "", 1)
			} else if seperated[0] == "CONDJUMP" {
				return "condjump|" + strings.Replace(seperated[1], ")", "", 1)
			} else if seperated[0] == "SLEEP" {
				return "sleep|" + strings.Replace(seperated[1], ")", "", 1)
			} else if command == "LEN" {
				return "length"
			} else {
				return "Push|" + command
			}
		} else {
			return "Push|" + command
		}
	} else {
		return "Push|" + ("string:" + strings.Split(command, "|")[0])
	}

}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func TestCondition(condition string, topValue int) bool {
	operator := condition[0]

	strValue := trimFirstRune(condition)

	value, err := strconv.ParseInt(strValue, 10, 64)

	if err != nil {
		log.Fatal("bad condition")
	}

	switch operator {
	case ('>'):
		if topValue > int(value) {
			return true
		} else {
			return false
		}
	case ('<'):
		if topValue < int(value) {
			return true
		} else {
			return false
		}
	case '=':
		if topValue == int(value) {
			return true
		} else {
			return false
		}
	case '!':
		if topValue != int(value) {
			return true
		} else {
			return false
		}
	default:
		log.Fatal("Bad operator")
		return false
	}
}

func CheckString(command string) (string, bool) {
	var seperated = strings.Split(command, ":")

	if len(seperated) != 1 {
		return seperated[1], true
	} else {
		return "", false
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
