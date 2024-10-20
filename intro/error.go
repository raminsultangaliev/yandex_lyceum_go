package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func ConcatStringsAndInt(str1, str2 string, num int) string {
	return str1 + " " + str2 + " " + fmt.Sprintf("%d", num)
}

var (
	ErrDivisionByZero          = errors.New("division by zero is not allowed")
	ErrPositionOutOfRange      = errors.New("position out of range")
	ErrNegativeNumberFactorial = errors.New("factorial is not defined for negative numbers")
	ErrNegativeNumber          = errors.New("negative numbers are not allowed")
	ErrNotIntegers             = errors.New("invalid input, please provide two integers")
)

func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return float64(a) / float64(b), nil
}

func GetCharacterAtPosition(str string, position int) (rune, error) {
	s := []rune(str)
	if position >= len(s) {
		return rune(0), ErrPositionOutOfRange
	}
	return s[position], nil
}

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativeNumberFactorial
	}

	if n == 0 {
		return 1, nil
	}

	prev, ok := Factorial(n - 1)

	if ok != nil {
		return 0, ok
	}

	return prev * n, nil
}

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", ErrNegativeNumber
	}

	res := ""
	for num > 0 {
		res = fmt.Sprintf("%b", num%2) + res
		num /= 2
	}

	return res, nil
}

func StrToInt(str string) (int, error) {
	res := 0
	var start int
	var is_negative bool

	if str[0] == '-' && len(str) > 1 {
		start = 1
		is_negative = true
	} else {
		start = 0
		is_negative = false
	}

	for _, i := range str[start:] {
		if unicode.IsDigit(i) {
			res = res*10 + int(i) - int('0')
		} else {
			return 0, ErrNotIntegers
		}
	}

	if is_negative {
		res *= -1
	}

	return res, nil
}

func SumTwoIntegers(a, b string) (int, error) {
	first, err1 := StrToInt(a)
	second, err2 := StrToInt(b)

	if err1 != nil || err2 != nil {
		return 0, ErrNotIntegers
	}

	return first + second, nil
}

func AreAnagrams(str1, str2 string) bool {
	m1, m2 := make(map[rune]int), make(map[rune]int)

	for _, r := range strings.ToLower(str1) {
		m1[r]++
	}

	for _, r := range strings.ToLower(str2) {
		m2[r]++
	}

	for k, _ := range m1 {
		if m2[k] != m1[k] {
			return false
		}
	}

	for k, _ := range m2 {
		if m2[k] != m1[k] {
			return false
		}
	}

	return true
}

func main() {
	str := "-"
	fmt.Println(StrToInt(str))
	fmt.Println(SumTwoIntegers("123", "6123"))
	fmt.Println(SumTwoIntegers("-50", "2"))
	fmt.Println(AreAnagrams("a", "A"))
}
