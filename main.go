package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var currentline = 0
var currentFile *os.File
var currentLines []string
var scanCounter = 0
var accValue string

func setUpDispatch(s *Stack) {
	s.dispatch = make(map[string]func(command string))

	s.dispatch["POP"] = s.POP
	s.dispatch["PUSH"] = func(command string) { s.Push(accValue) }
	s.dispatch["+"] = s.Add
	s.dispatch["-"] = s.Subtract
	s.dispatch["*"] = s.Multiply
	s.dispatch["/"] = s.Divide
	s.dispatch["^"] = s.Exponentiate
	s.dispatch["DUPE"] = s.Dupe
	s.dispatch["STDIN"] = s.STDIN
	s.dispatch["INPUT"] = s.INPUT
	s.dispatch["JUMP"] = s.JUMP
	s.dispatch["CONDJUMP"] = s.CONDJUMP
	s.dispatch["LENGTH"] = s.LENGTH
	s.dispatch["SLEEP"] = s.SLEEP
	s.dispatch["PRINT"] = s.PRINT
	s.dispatch["PRINTSTR"] = s.PRINTSTR
	s.dispatch["CLEAR"] = s.CLEAR
	s.dispatch["CONCAT"] = s.CONCAT
	s.dispatch["LOAD"] = s.LOAD
	s.dispatch["SCAN"] = s.SCAN
	s.dispatch["SCANLN"] = s.SCANLN
	s.dispatch["UNLOAD"] = s.UNLOAD
	s.dispatch["RUN"] = s.RUN
	s.dispatch["RUNFILE"] = s.RUNFILE
	s.dispatch["SCANFILE"] = s.SCANFILE
}

func main() {
	var f []byte
	var err error
	if len(os.Args) == 1 {
		f, err = ioutil.ReadFile("main.txt")
	} else {
		f, err = ioutil.ReadFile(os.Args[1])
	}

	if err != nil {
		log.Fatalf("No file %s exists. Try adding a complete path", "stack.txt")
	}

	var s Stack

	setUpDispatch(&s)

	var lines = strings.Split(string(f), ";")

	for i, line := range lines {
		line = strings.ReplaceAll(line, "\n", "")
		if line != "" {
			if line[0] == ' ' {
				line = trimFirstRune(line)
			}
			lines[i] = line
		} else {
			remove(lines, i)
		}

	}

	for _ = 0; currentline < len(lines); currentline++ {
		var cmd = lines[currentline]
		var seperated = strings.Split(cmd, "(")
		var checkLoad = strings.Split(cmd, "<>")

		if len(checkLoad) != 0 {
			if len(seperated) == 0 {
				var stackCommand = s.dispatch[cmd]
				if stackCommand != nil {
					stackCommand(cmd)
				} else {
					s.Push(cmd)
				}
			} else {
				var stackCommand = s.dispatch[seperated[0]]
				if stackCommand != nil {
					stackCommand(cmd)
				} else {
					s.Push(cmd)
				}
			}
		} else {
			if checkLoad[0] == "LOAD" {
				s.LOAD(checkLoad[1])
			}
		}

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

type Stack struct {
	stack    []string
	dispatch map[string]func(command string)
}

func (s *Stack) Push(command string) {
	s.stack = append(s.stack, command)

}

func (s *Stack) Pop(command string) string {
	if len(s.stack) == 0 {
		log.Fatalf("No items in stack to pop")
	}

	var poppedStack []string
	var lastElement = s.stack[len(s.stack)-1]

	poppedStack = s.stack[:len(s.stack)-1]

	s.stack = poppedStack

	return lastElement
}

func (s *Stack) Dupe(command string) {
	cmd := s.Pop(command)
	s.Push(cmd)
	s.Push(cmd)
}

func (s *Stack) Add(command string) {
	rawFirstNumber := s.Pop(command)
	rawSecondNumber := s.Pop(command)
	firstNumber, error1 := strconv.ParseInt(rawFirstNumber, 10, 64)
	secondNumber, error2 := strconv.ParseInt(rawSecondNumber, 10, 64)
	s.Push(rawSecondNumber)
	s.Push(rawFirstNumber)
	if error1 != nil || error2 != nil {
		log.Fatal("Invalid Addition")
	} else {
		s.Push(strconv.FormatInt(secondNumber+firstNumber, 10))
	}
}

func (s *Stack) Subtract(command string) {
	rawFirstNumber := s.Pop(command)
	rawSecondNumber := s.Pop(command)
	firstNumber, error1 := strconv.ParseInt(rawFirstNumber, 10, 64)
	secondNumber, error2 := strconv.ParseInt(rawSecondNumber, 10, 64)
	s.Push(rawSecondNumber)
	s.Push(rawFirstNumber)
	if error1 != nil || error2 != nil {
		log.Fatal("Invalid Subtraction")
	} else {
		s.Push(strconv.FormatInt(secondNumber-firstNumber, 10))
	}
}

func (s *Stack) Multiply(command string) {
	rawFirstNumber := s.Pop(command)
	rawSecondNumber := s.Pop(command)
	firstNumber, error1 := strconv.ParseInt(rawFirstNumber, 10, 64)
	secondNumber, error2 := strconv.ParseInt(rawSecondNumber, 10, 64)
	s.Push(rawSecondNumber)
	s.Push(rawFirstNumber)
	if error1 != nil || error2 != nil {
		log.Fatal("Invalid Multiplication")
	} else {
		s.Push(strconv.FormatInt(secondNumber*firstNumber, 10))
	}
}

func (s *Stack) Divide(command string) {
	rawFirstNumber := s.Pop(command)
	rawSecondNumber := s.Pop(command)
	firstNumber, error1 := strconv.ParseInt(rawFirstNumber, 10, 64)
	secondNumber, error2 := strconv.ParseInt(rawSecondNumber, 10, 64)
	s.Push(rawSecondNumber)
	s.Push(rawFirstNumber)
	if error1 != nil || error2 != nil {
		log.Fatal("Invalid Division")
	} else {
		s.Push(strconv.FormatInt(secondNumber/firstNumber, 10))
	}
}

func (s *Stack) Exponentiate(command string) {
	rawFirstNumber := s.Pop(command)
	rawSecondNumber := s.Pop(command)
	firstNumber, error1 := strconv.ParseInt(rawFirstNumber, 10, 64)
	secondNumber, error2 := strconv.ParseInt(rawSecondNumber, 10, 64)
	s.Push(rawSecondNumber)
	s.Push(rawFirstNumber)
	if error1 != nil || error2 != nil {
		log.Fatal("Invalid Exponentiation")
	} else {
		s.Push(strconv.FormatInt(int64(math.Pow(float64(secondNumber), float64(firstNumber))), 10))
	}
}

func (s *Stack) STDIN(command string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s.Push(scanner.Text())
	}
}

func (s *Stack) INPUT(command string) {
	seperated := strings.Split(command, "(")

	if len(seperated) != 2 {
		log.Fatal("Invalid input, add parentheses")
	} else {
		var input string
		fmt.Print(strings.Replace(seperated[1], ")", "", 1))
		fmt.Scanln(&input)
		s.Push(input)
	}
}

func (s *Stack) PRINT(command string) {
	topCmd := s.Pop(command)
	fmt.Println(topCmd)
	s.Push(topCmd)
}

func (s *Stack) PRINTSTR(command string) {
	topCmd := s.Pop(command)
	var seperated = strings.Split(topCmd, ":")
	println(seperated[1])
	s.Push(topCmd)
}

func (s *Stack) CLEAR(command string) {
	for i := 0; i < len(s.stack); i++ {
		s.Pop(command)
	}
}

func (s *Stack) JUMP(command string) {
	seperated := strings.Split(command, "(")

	if len(seperated) != 2 {
		log.Fatal("Invalid jump, add parentheses")
	} else {
		usableCmd := strings.Replace(seperated[1], ")", "", 1)

		if usableCmd == "__TOPREPLACE__" {
			usableCmd = s.Pop(command)
			s.Push(usableCmd)
		} else if usableCmd == "__TOP__" {
			usableCmd = s.Pop(command)
		}
		tryInt, err := strconv.ParseInt(usableCmd, 10, 64)

		if err != nil {
			println(usableCmd)
			log.Fatal("invalid input to jump")
		} else {
			currentline = int(tryInt - 1)
		}
	}
}

func (s *Stack) CONDJUMP(command string) {
	seperated := strings.Split(command, "(")

	condition := strings.Split(seperated[1], ":")

	topValStr := s.Pop(command)
	s.Push(topValStr)

	topVal, err := strconv.ParseInt(topValStr, 10, 64)

	if err != nil {
		log.Fatal("invalid top value")
	}

	if len(condition) != 2 {
		log.Fatal("invalid condition")
	} else {
		if condition[0] == "LINES" {
			if scanCounter < len(currentLines) {
				if len(seperated) != 2 {
					log.Fatal("Invalid jump, add parentheses")
				} else {
					usableCmd := strings.Replace(condition[1], ")", "", 1)

					if usableCmd == "__TOPREPLACE__" {
						usableCmd = s.Pop(command)
						s.Push(usableCmd)
					} else if usableCmd == "__TOP__" {
						usableCmd = s.Pop(command)
					}
					tryInt, err := strconv.ParseInt(usableCmd, 10, 64)

					if err != nil {
						log.Fatal("invalid input to jump")
					} else {
						currentline = int(tryInt - 1)
					}
				}
			}
		} else {
			if TestCondition(condition[0], int(topVal)) {
				if len(seperated) != 2 {
					log.Fatal("Invalid jump, add parentheses")
				} else {
					usableCmd := strings.Replace(condition[1], ")", "", 1)

					if usableCmd == "__TOPREPLACE__" {
						usableCmd = s.Pop(command)
						s.Push(usableCmd)
					} else if usableCmd == "__TOP__" {
						usableCmd = s.Pop(command)
					}
					tryInt, err := strconv.ParseInt(usableCmd, 10, 64)

					if err != nil {
						log.Fatal("invalid input to jump")
					} else {
						currentline = int(tryInt - 1)
					}
				}
			}
		}
	}
}

func (s *Stack) SLEEP(command string) {
	seperated := strings.Split(command, "(")

	if len(seperated) != 2 {
		log.Fatal("Invalid sleep, add parentheses")
	} else if strings.ReplaceAll(seperated[1], ")", "") == "__ACC__" {
		cmd := accValue
		tryInt, err := strconv.ParseInt(cmd, 10, 64)

		if err != nil {
			log.Fatal("invalid input to sleep")
		} else {
			time.Sleep(time.Duration(tryInt) * time.Second)
		}
	} else {
		usableCmd := strings.Replace(seperated[1], ")", "", 1)
		tryInt, err := strconv.ParseInt(usableCmd, 10, 64)

		if err != nil {
			log.Fatal("invalid input to sleep")
		} else {
			time.Sleep(time.Duration(tryInt) * time.Second)
		}
	}
}

func (s *Stack) LENGTH(command string) {
	s.Push(strconv.FormatInt(int64(len(s.stack)), 10))
}

func (s *Stack) CONCAT(command string) {
	firstString := s.Pop(command)
	secondString := s.Pop(command)
	s.Push(secondString + firstString)
}

func (s *Stack) LOAD(command string) {

	seperated := strings.Split(command, "(")

	if len(seperated) != 2 {
		log.Fatal("Invalid load, add parentheses")
	} else if seperated[1] == ")" {
		cmd := s.Pop(command)
		s.Push(cmd)
		usableCmd := strings.Split(cmd, ":")[0]
		f, err := os.Open(usableCmd)
		if err != nil {
			log.Fatal("Error opening file")
		} else {
			currentFile = f
		}
	} else if strings.ReplaceAll(seperated[1], ")", "") == "__ACC__" {
		var usableValue string
		if len(strings.Split(accValue, "\"")) != 0 {
			usableValue = strings.Split(accValue, "\"")[0]
		}
		f, err := os.Open(usableValue)
		if err != nil {
			log.Fatal("Error opening file")
		} else {
			currentFile = f
		}
	} else {
		usableCmd := strings.Replace(seperated[1], ")", "", 1)
		f, err := os.Open(usableCmd)
		if err != nil {
			log.Fatal("Error opening file")
		} else {
			currentFile = f
		}
	}
}

func (s *Stack) SCAN(command string) {
	if currentFile == nil || currentLines != nil {
		log.Fatal("invalid scan")
	} else {
		scanner := bufio.NewScanner(currentFile)
		for i := 0; scanner.Scan(); i++ {
			currentLines = append(currentLines, scanner.Text())
		}
	}
}

func (s *Stack) SCANLN(command string) {
	if scanCounter < len(currentLines) {
		s.Push(currentLines[scanCounter])
	}
	scanCounter++
}

func (s *Stack) UNLOAD(command string) {
	currentFile = nil
	currentLines = nil
	scanCounter = 0
}

func (s *Stack) RUN(command string) {
	var cmd = s.Pop("")

	var stackCommand = s.dispatch[cmd]
	if stackCommand != nil {
		stackCommand(cmd)

	} else {
		s.Push(cmd)
	}
}

func (s *Stack) RUNFILE(command string) {
	s.LOAD(command)
	s.SCAN(command)
	for i := 0; i < len(currentLines); i++ {

		s.SCANLN(command)
		s.RUN(command)
	}
	s.UNLOAD(command)
}

func (s *Stack) SCANFILE(command string) {
	s.LOAD(command)
	s.SCAN(command)
	for i := 0; i < len(currentLines); i++ {
		s.SCANLN(command)
	}
	s.UNLOAD(command)
}

func (s *Stack) POP(command string) {
	if len(s.stack) == 0 {
		log.Fatalf("No items in stack to pop")
	}

	var poppedStack []string
	var lastElement = s.stack[len(s.stack)-1]

	poppedStack = s.stack[:len(s.stack)-1]

	s.stack = poppedStack

	accValue = lastElement
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
