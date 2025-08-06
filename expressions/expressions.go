// Package expressions contains a simple expression evaluator.
package expressions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type funcName uint8

const (
	funcUnknown funcName = iota
	funcAdd
	funcSubtract
	funcMultiply
	funcDivide
	funcSqrt
	funcSin
)

type functionClass interface {
	function() funcName
}

type binaryFuncClass struct {
	funcName
}

func (b *binaryFuncClass) function() funcName {
	return b.funcName
}

type unaryFuncClass struct {
	funcName
}

func (u *unaryFuncClass) function() funcName {
	return u.funcName
}

type Env map[rune]float64

func isDigit(x rune) bool {
	return '0' <= x && x <= '9'
}

func getFunc(name string) (functionClass, error) {
	switch name {
	case "add":
		return &binaryFuncClass{funcAdd}, nil
	case "subtract":
		return &binaryFuncClass{funcSubtract}, nil
	case "multiply":
		return &binaryFuncClass{funcMultiply}, nil
	case "divide":
		return &binaryFuncClass{funcDivide}, nil
	case "sqrt":
		return &unaryFuncClass{funcSqrt}, nil
	case "sin":
		return &unaryFuncClass{funcSin}, nil
	default:
		return nil, fmt.Errorf("unknown function name: %v", name)
	}
}

func getArgs(argString string) ([]Constant, error) {
	var args []Constant
	for _, char := range strings.Split(argString, ",") {
		value, _ := strconv.ParseFloat(char, 64)
		args = append(args, Constant(value))
	}
	return args, nil
}

func Parse(inputString string) (Expr, error) {
	var (
		funcClass functionClass
		args      []Constant
		err       error
	)
	processedString := strings.ToLower(strings.ReplaceAll(inputString, ", ", ","))
	l := 0
	for r, char := range processedString {
		if char == '(' {
			funcClass, err = getFunc(processedString[l:r])
			if err != nil {
				return nil, fmt.Errorf("unknown function name: %v", processedString[l:r])
			}
			l = r + 1
		} else if char == ')' {
			args, err = getArgs(processedString[l:r])
			if err != nil {
				return nil, fmt.Errorf("received invalid arguments: %v", args)
			}
			l = r + 1
		}
	}
	return check(funcClass, args)
}

func check(funcClass functionClass, args []Constant) (Expr, error) {
	switch funcClass.(type) {
	case *binaryFuncClass:
		if len(args) != 2 {
			return nil, fmt.Errorf("binary expressions require 2 inputs, received: %v", len(args))
		}
		left, right := args[0], args[1]
		return &binaryExpr{funcClass.function(), left, right}, nil
	case *unaryFuncClass:
		if len(args) != 1 {
			return nil, fmt.Errorf("unary expression require 1 input, received: %v", len(args))
		}
		child := args[0]
		return &unaryExpr{funcClass.function(), child}, nil
	default:
		return nil, fmt.Errorf("unknown function class: %v", funcClass)
	}
}

type Expr interface {
	Evaluate() (float64, error)
}

type Constant float64

func (c Constant) Evaluate() (float64, error) {
	return float64(c), nil
}

type binaryExpr struct {
	funcName
	left, right Expr
}

func (b *binaryExpr) Evaluate() (float64, error) {
	left, _ := b.left.Evaluate()
	right, _ := b.right.Evaluate()
	switch b.funcName {
	case funcAdd:
		return left + right, nil
	case funcSubtract:
		return left - right, nil
	case funcMultiply:
		return left * right, nil
	case funcDivide:
		return left / right, nil
	default:
		return 0, fmt.Errorf("unknown function expression: %v", b.funcName)
	}
}

type unaryExpr struct {
	funcName
	child Expr
}

func (u *unaryExpr) Evaluate() (float64, error) {
	child, _ := u.child.Evaluate()
	switch u.funcName {
	case funcSqrt:
		return math.Sqrt(child), nil
	case funcSin:
		return math.Sin(child), nil
	default:
		return 0, fmt.Errorf("unknown function expression: %v", u.funcName)
	}
}
