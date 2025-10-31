package main

import "fmt"

// There is no inheritance in golang
// but it is possible to achieve with COMPOSITION

type Person struct {
	Name string
	Age  int
}

func (p *Person) Intro() {
	fmt.Println("Name:", p.Name, "| Age:", p.Age)
}

type Employee struct {
	*Person
	EmployeeId int
}

func (e *Employee) Intro() {
	fmt.Println("Name:", e.Name, "| Age:", e.Age, "| EmployeeID:", e.EmployeeId)
}

func main() {

	p1 := Person{"Atharva", 24}
	e1 := Employee{&Person{"Test", 19}, 123332}
	e2 := Employee{&p1, 1111}

	p1.Intro()
	e1.Intro()
	e2.Intro()
}
