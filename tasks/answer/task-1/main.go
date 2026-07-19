package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")

	name := "Bharathi"
	goal := "Become Devops Engineer"
	city := "Madurai"
	now := time.Now()

	fmt.Println("Name: ", name)
	fmt.Println("I am from ", city)
	fmt.Println("My learning goal is ", goal)
	fmt.Println("Date now is ", now.Format("2006-01-02"))

}
