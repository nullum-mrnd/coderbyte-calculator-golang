package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func calculator(str string) int {
	var splitted = strings.Split(str, "")
	var operatorsQueue []string
	var outputQueue []int
	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	ops := []string{"+", "-", "*", "/"}
	currentSize := len(splitted)

	for i := 0; i < currentSize; i++ {
		if slices.Contains(nums, string(splitted[i])) {
			n, add := collectNumber(splitted[i:], nums)
			i += add
			outputQueue = append(outputQueue, n)
			continue
		}
		if slices.Contains(ops, string(splitted[i])) {
			if string(splitted[i]) == "/" || string(splitted[i]) == "*" {
				var a int
				var op string
				outputQueue, a = popInt(outputQueue)
				if splitted[i] == "*" && splitted[i+1] == "*" {
					op = "**"
					i++
				} else {
					op = string(splitted[i])
				}
				b, add := collectNumber(splitted[i+1:], nums)
				r := calculate(a, b, op)
				outputQueue = append(outputQueue, r)
				i += add + 1
				continue
			}
			operatorsQueue = append(operatorsQueue, string(splitted[i]))
		}
		if splitted[i] == "(" {
			b := matchParenthesis(splitted[i+1:])
			r := calculator(strings.Join(splitted[i+1:i+b], ""))
			splitted[i] = fmt.Sprint(r)
			splitted = slices.Delete(splitted, i+1, i+b+1)
			outputQueue = append(outputQueue, r)
			currentSize = len(splitted)
		}
	}

	for len(outputQueue) > 1 {
		var a int
		var b int
		var op string
		a = outputQueue[0]
		b = outputQueue[1]
		op = operatorsQueue[0]
		operatorsQueue = operatorsQueue[1:]
		r := calculate(a, b, op)
		if r == 0 && operatorsQueue[len(operatorsQueue)-1] == "-" {
			operatorsQueue, _ = popStr(operatorsQueue)
		}
		outputQueue[1] = r
		outputQueue = outputQueue[1:]
	}
	return outputQueue[0]
}

//Auxiliary functions

func toInt(str string) int {
	b, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Erro: ", err)
		os.Exit(1)
	}
	return b
}

func calculate(a int, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "**":
		return int(math.Pow(float64(a), float64(b)))
	}
	return -99999
}

func popInt(s []int) ([]int, int) {
	if len(s) == 0 {
		return s, 0
	}
	return s[:len(s)-1], s[len(s)-1]
}
func popStr(s []string) ([]string, string) {
	if len(s) == 0 {
		return s, ""
	}
	return s[:len(s)-1], s[len(s)-1]
}

func remove(s []string) ([]string, string) {
	return s[1:], s[0]
}

func collectNumber(splitted []string, nums []string) (int, int) {
	i := 0
	var strResult string
	for i < len(splitted) && slices.Contains(nums, splitted[i]) {
		strResult += splitted[i]
		i++
	}
	r, err := strconv.Atoi(strResult)
	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}

	return r, i - 1
}

func matchParenthesis(s []string) int {
	level := 0
	for i, v := range s {
		if v == ")" {
			if level == 0 {
				return i + 1
			} else {
				level--
			}
		}
		if v == "(" {
			level += 1
		}
	}
	return 99999
}
