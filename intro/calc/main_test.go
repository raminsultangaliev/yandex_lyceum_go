package main

import (
	"testing"
)

func TestCalcSyntaxError(t *testing.T) {
	cases := []string {"a", "2+(", "1+", "1*", "()", "+", "(+)"}

	for _, tc := range cases {
		t.Run(tc, func(t *testing.T) {
			_, err := Calc(tc)
			if err != ErrSyntax {
				t.Errorf("Calc(%q) = %v; want %v", tc, err, ErrSyntax)
			}
		})
	}
}

func TestCalc(t *testing.T) {
	cases := []struct {
		in  string
		out float64
		err error
	}{
		{"1+2+3+4+5", 15, nil},					   // Simple addition
		{"1/2", 0.5, nil},						   // Division
		{"1/0", 0, ErrDivisionByZero},			   // Division by zero
		{"(10*9)+(2/5)+2*3*10", 150.4, nil},	   // Random test
		{"5-3+2", 4, nil},                         // Mixed operations
		{"6*7", 42, nil},                          // Simple multiplication
		{"10/(5-5)", 0, ErrDivisionByZero},        // Division by zero after computation
		{"2*(3+4)", 14, nil},                      // Testing brackets for priority
		{"(2.5+2.5)*2", 10, nil},                  // Floating point addition and multiplication
		{"10-3*2", 4, nil},                        // Testing operator precedence
		{"(5+5)/(2+3)", 2, nil},                   // Mixed operations with brackets
		{"3.5*2", 7, nil},                         // Simple floating point multiplication
		{"(2+3)*(1.5+1.5)", 15, nil},              // Combination of integers and floats with brackets
		{"((2+3)*2)-((1+1)*2)", 6, nil},           // Nested brackets
		{"4*(3+(2-1))", 16, nil},                  // Nested operations within brackets
		{"4.2-1.2", 3.0, nil},                     // Simple floating point subtraction
		{"1-2-3-4", -8, nil},                      // Chain of subtraction
		{"((1+1))", 2, nil},                       // Double brackets
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			got, err := Calc(tc.in)
			if got != tc.out || err != tc.err {
				t.Errorf("Calc(%q) = %f, %v; want %f, %v", tc.in, got, err, tc.out, tc.err)
			}
		})
	}
}
