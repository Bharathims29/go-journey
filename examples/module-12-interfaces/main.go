package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(message string) {
	fmt.Println("LOG:", message)
}

func printArea(shape Shape) {
	fmt.Printf("Area: %.2f\n", shape.Area())
}

func processPayment(amount float64, logger Logger) {
	logger.Log(fmt.Sprintf("processed payment: %.2f", amount))
}

func describe(value any) {
	switch v := value.(type) {
	case string:
		fmt.Println("string value:", v)
	case int:
		fmt.Println("int value:", v)
	case bool:
		fmt.Println("bool value:", v)
	default:
		fmt.Printf("unknown type: %T\n", v)
	}
}

func main() {
	fmt.Println("Module 12: Interfaces")
	fmt.Println()

	fmt.Println("Shape interface")
	printArea(Rectangle{Width: 10, Height: 5})
	printArea(Circle{Radius: 3})
	fmt.Println()

	fmt.Println("Logger interface")
	processPayment(499.99, ConsoleLogger{})
	fmt.Println()

	fmt.Println("Type assertion")
	var value any = "learning interfaces"
	text, ok := value.(string)
	if ok {
		fmt.Println("asserted string:", text)
	}
	fmt.Println()

	fmt.Println("Type switch")
	describe("go")
	describe(42)
	describe(true)
	describe(3.14)
}
