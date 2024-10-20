package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	d := b*b - 4.0*a*c
	if d < 0 {
		fmt.Println(0, 0)
		return
	}
	x1, x2 := (-b-math.Sqrt(d))/(2.0*a), (-b+math.Sqrt(d))/(2.0*a)
	if d == 0 {
		fmt.Println(x1)
	} else {
		fmt.Println(x1, x2)
	}
}
