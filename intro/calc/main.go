package main

import (
	"errors"
	"strconv"
)

var (
	ErrSyntax         = errors.New("syntax error")
	ErrDivisionByZero = errors.New("division by zero is not allowed")
)

func isAllowedCharacter(r rune) bool {
	allowed := []rune{'+', '-', '*', '/', '(', ')', '.'}
	for _, c := range allowed {
		if r == c {
			return true
		}
	}
	return (r >= '0' && r <= '9')
}

func FindOpsAndExps(expression string, mode int) ([]rune, []string, error) {
	depth := 0
	from := 0
	ops := []rune{}
	exps := []string{}

	for i, r := range expression {
		if !isAllowedCharacter(r) || depth < 0 {
			return ops, exps, ErrSyntax
		}

		if r == '(' {
			depth += 1
		} else if r == ')' {
			depth -= 1
		}
		
		if depth == 0 && (mode == 0 && (r == '+' || r == '-') || mode == 1 && (r == '*' || r == '/')) {
			ops = append(ops, r)
			exps = append(exps, expression[from : i])
			if exps[len(exps) - 1] == "" {
				return ops, exps, ErrSyntax
			}
			from = i + 1
		}
	}

	exps = append(exps, expression[from :])
	if exps[len(exps) - 1] == "" || depth != 0 {
		return ops, exps, ErrSyntax
	}

	return ops, exps, nil
}

func ExpToFloat(exp string) (float64, error) {
	var res float64
	var err error
	if exp[0] == '(' {
		if len(exp) == 2 {
			return 0, ErrSyntax
		}
		res, err = Calc(exp[1 : len(exp) - 1]) 
	} else {
		res, err = Calc(exp)
	}
	if err != nil {
		return 0, err
	}
	return res, nil
}

func ApplyOps(ops []rune, exps []string) (float64, error) {
	var err error
	res, err := ExpToFloat(exps[0])
	if err != nil {
		return 0, err
	}

	for i, op := range ops {
		temp, err := ExpToFloat(exps[i + 1])
		if err != nil {
			return 0, err
		}

		switch op {
		case '+':
			res += temp
		case '-':
			res -= temp
		case '*':
			res *= temp
		case '/':
			if temp == 0 {
				return 0, ErrDivisionByZero
			}
			res /= temp
		}
	}

	return res, nil
}

func Calc(expression string) (float64, error) {
	if expression == "" || expression == "()" {
		return 0, ErrSyntax
	}

	var err error
	lowOps, lowExps, err := FindOpsAndExps(expression, 0)
	if err != nil {
		return 0, err
	}

	if len(lowOps) == 0 {
		highOps, highExps, err := FindOpsAndExps(lowExps[0], 1)
		if err != nil {
			return 0, err
		}
		
		if len(highOps) == 0 {
			var f float64
			var err error
			if highExps[0][0] == '(' && highExps[0][len(highExps[0]) - 1] == ')' {
				f, err = Calc(highExps[0][1 : len(highExps[0]) - 1])
			} else {
				f, err = strconv.ParseFloat(highExps[0], 64)
			}
			return f, err
		} else {
			return ApplyOps(highOps, highExps)
		}
	}

	return ApplyOps(lowOps, lowExps)
} 	