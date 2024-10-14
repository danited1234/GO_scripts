package main

import (
	"fmt"
)

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}
type fullTime struct {
	name   string
	salary int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

func (f fullTime) getSalary() int {
	return f.salary
}
func (f fullTime) getName() string {
	return f.name
}
func test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
}
func main() {
	test(fullTime{name: "Ahmad Hassan", salary: 5000})
	test(contractor{name: "Aladdin", hourlyPay: 587, hoursPerYear: 7884})
}
