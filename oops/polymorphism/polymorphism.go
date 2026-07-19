package polymorphism

import "fmt"

type Shape interface {
	Render()
}

type Circle struct{}

func (C Circle) Render() {
	fmt.Println(" Its a circle")
}

type Square struct{}

func (S Square) Render() {
	fmt.Println("Its a Square")
}
