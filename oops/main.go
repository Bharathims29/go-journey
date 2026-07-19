package main

import (
	"fmt"
	"oops/composition"
	"oops/person"
	"oops/polymorphism"
)

func main() {
	fmt.Println("---------------Method 1----------------")
	p1 := person.New("Bharathi", "MS", 25)
	p2 := person.New("Luffy", "Monkey", 30)

	fmt.Println(p1.FirstName())
	fmt.Println(p2.FirstName())

	fmt.Println(p1.Age())
	fmt.Println(p2.Age())

	fmt.Println(p1.FamilyName())
	fmt.Println(p2.FamilyName())
	fmt.Println("--------------Method 2-----------------")
	p := person.Person{}
	p.SetFirstName("Bharathi")
	p.SetLastName("MS")
	p.SetAge(25)
	fmt.Println(p.FirstName() + " " + p.LastName() + " age is " + p.Age())

	fmt.Println("--------------Polymorphism-----------------")
	var c polymorphism.Shape = polymorphism.Circle{}
	var d polymorphism.Shape = polymorphism.Square{}
	c.Render()
	d.Render()

	fmt.Println("--------------Polymorphism-----------------")

	car1 := composition.Newcar("audi", 1200, 123)
	fmt.Println(car1.Engine.HP())
	bike1 := composition.Newbike("BMW", 1100, 80)
	fmt.Println(bike1.HP())

}
