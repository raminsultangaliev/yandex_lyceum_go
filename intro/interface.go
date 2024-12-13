package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

func CalculateArea(shape Shape) float64 {
	return shape.Area()
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Animal interface {
	MakeSound()
}

type Cat struct{}

type Dog struct{}

func (d Dog) MakeSound() {
	fmt.Println("Гав!")
}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}

type Logger interface {
	Log()
}

type Log struct {
	level LogLevel
}

func (l Log) Log(message string) {
	switch l.level {
	case Error:
		fmt.Printf("ERROR: %s\n", message)
	case Info:
		fmt.Printf("INFO: %s\n", message)
	}

}

type LogLevel string

const Error LogLevel = "Error"
const Info LogLevel = "Info"

// func main() {
// 	r := Rectangle{width: 10, height: 20}
// 	c := Circle{radius: 6}
// 	fmt.Println(CalculateArea(r))
// 	fmt.Println(CalculateArea(c))
// }
