package rpn_test

import (
	"testing"
	rpn "github.com/raminsultangaliev/rpn/pkg/rpn"
)

func TestCalcSyntaxError(t *testing.T) {
	cases := []string {"a", "2+(", "1+", "1*", "()", "+", "(+)"}

	for _, tc := range cases {
		t.Run(tc, func(t *testing.T) {
			_, err := rpn.Calc(tc)
			if err != rpn.ErrSyntax {
				t.Errorf("Calc(%q) = %v; want %v", tc, err, rpn.ErrSyntax)
			}
		})
	}
}

func TestCalc(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  float64
		err  error
	}{
		{"Simple addition", "1+2+3+4+5", 15, nil},
		{"Division", "1/2", 0.5, nil},
		{"Division by zero", "1/0", 0, rpn.ErrDivisionByZero},
		{"Random test", "(10*9)+(2/5)+2*3*10", 150.4, nil},
		{"Mixed operations", "5-3+2", 4, nil},
		{"Simple multiplication", "6*7", 42, nil},
		{"Division by zero after computation", "10/(5-5)", 0, rpn.ErrDivisionByZero},
		{"Testing brackets for priority", "2*(3+4)", 14, nil},
		{"Floating point addition and multiplication", "(2.5+2.5)*2", 10, nil},
		{"Testing operator precedence", "10-3*2", 4, nil},
		{"Mixed operations with brackets", "(5+5)/(2+3)", 2, nil},
		{"Simple floating point multiplication", "3.5*2", 7, nil},
		{"Combination of integers and floats with brackets", "(2+3)*(1.5+1.5)", 15, nil},
		{"Nested brackets", "((2+3)*2)-((1+1)*2)", 6, nil},
		{"Nested operations within brackets", "4*(3+(2-1))", 16, nil},
		{"Simple floating point subtraction", "4.2-1.2", 3, nil},
		{"Chain of subtraction", "1-2-3-4", -8, nil},
		{"Double brackets", "((1+1))", 2, nil},
		{"Infinite fraction", "1/(0.1+0.2)", 3.333333333333333, nil},
		// {"Negative number", "-5", -5, nil},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			got, err := rpn.Calc(tc.in)
			if got != tc.out || err != tc.err {
				t.Errorf("Calc(%q) = %f, %v; want %f, %v", tc.in, got, err, tc.out, tc.err)
			}
		})
	}
}