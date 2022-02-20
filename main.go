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

var currentline = 0
var currentFile *os.File
var currentLines []string
var scanCounter = 0

func main() {
	//fScanner := bufio.NewScanner(os.Stdin)

	//fScanner.Scan()
	//filename := fScanner.Text()

	f, err := os.Open("stack.txt")

	if err != nil {
		//log.Fatalf("No file %s exists. Try adding a complete path", filename)
	}

	var currentValue string

	var s Stack

	s.functionMap = make(map[string]func(command string))

	s.functionMap["POP"] = func(command string) { currentValue = s.Pop(command) }
	s.functionMap["PUSH"] = func(command string) { s.Push(currentValue) }
	s.functionMap["+"] = s.Add
	s.functionMap["-"] = s.Subtract
	s.functionMap["*"] = s.Multiply
	s.functionMap["/"] = s.Divide
	s.functionMap["^"] = s.Exponentiate
	s.functionMap["DUPE"] = s.Dupe
	s.functionMap["STDIN"] = s.STDIN
	s.functionMap["INPUT"] = s.INPUT
	s.functionMap["JUMP"] = s.JUMP
	s.functionMap["CONDJUMP"] = s.CONDJUMP
	s.functionMap["LENGTH"] = s.LENGTH
	s.functionMap["SLEEP"] = s.SLEEP
	s.functionMap["PRINT"] = s.PRINT
	s.functionMap["PRINTSTR"] = s.PRINTSTR
	s.functionMap["CLEAR"] = s.CLEAR
	s.functionMap["CONCAT"] = s.CONCAT
	s.functionMap["LOAD"] = s.LOAD
	s.functionMap["SCAN"] = s.SCAN
	s.functionMap["SCANLN"] = s.SCANLN
	s.functionMap["UNLOAD"] = s.UNLOAD

	var lines []string

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}

	for _ = 0; currentline < len(lines); currentline++ {
		var cmd = lines[currentline]
		var seperated = strings.Split(cmd, "(")
		var checkLoad = strings.Split(cmd, "<>")

		if len(checkLoad) != 0 {
			if len(seperated) == 0 {
				var stackCommand = s.functionMap[cmd]
				if stackCommand != nil {
					stackCommand(cmd)
				} else {
					s.Push(cmd)
				}
			} else {
				var stackCommand = s.functionMap[seperated[0]]
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

func CheckString(command string) (string, bool) {
	var seperated = strings.Split(command, ":")

	if len(seperated) != 1 {
		return seperated[1], true
	} else {
		return "", false
	}
}

type Stack struct {
	stack       []string
	functionMap map[string]func(command string)
}

func (s *Stack) Push(command string) {
	s.stack = append(s.stack, command)
}

func (s *Stack) Pop(command string) string {
	var poppedStack []string
	var lastElement = s.stack[len(s.stack)-1]

	for i := 0; i < len(s.stack)-1; i++ {
		poppedStack = append(poppedStack, s.stack[i])
	}
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

func (s *Stack) SLEEP(command string) {
	seperated := strings.Split(command, "(")

	if len(seperated) != 2 {
		log.Fatal("Invalid sleep, add parentheses")
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

func (s *Stack) LOAD(command string) {
	seperated := strings.Split(command, "(")

	if len(seperated) != 2 {
		log.Fatal("Invalid load, add parentheses")
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
	s.Push(currentLines[scanCounter])
	scanCounter++
}

func (s *Stack) UNLOAD(command string) {
	currentFile = nil
	currentLines = nil
	scanCounter = 0
}
