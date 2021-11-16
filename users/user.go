package user

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id   int
	Name string
}
type Employee struct {
	User
	Active bool
}

func (e *Employee) DesactiveEmployee(c Cachier) {
	c.desactive()
}
func (e *Employee) desactive() {
	e.Active = false
}

func (e *Employee) Calculate(items ...float32) float32 {
	if !e.Active {
		fmt.Println("user desactive")
		return 0
	}

	var sum float32
	for _, item := range items {
		sum += item
	}
	return sum * 1.15
}

func NewEmployee(name string) *Employee {
	return &Employee{
		User: User{
			Id:   rand.Intn(1000),
			Name: name,
		},
		Active: true,
	}
}

type Admin interface {
	DesactiveEmployee(c Cachier)
}

type Cachier interface {
	Calculate(items ...float32) float32
	desactive()
}
