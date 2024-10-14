package main

import "fmt"

// type employee interface {
// 	getName() string
// 	getSalary() int
// }

// type contractor struct {
// 	name         string
// 	hourlyPay    int
// 	hoursPerYear int
// }
// type fullTime struct {
// 	name   string
// 	salary int
// }

// func (c contractor) getName() string {
// 	return c.name
// }

// func (c contractor) getSalary() int {
// 	return c.hourlyPay * c.hoursPerYear
// }

// func (f fullTime) getSalary() int {
// 	return f.salary
// }
// func (f fullTime) getName() string {
// 	return f.name
// }
// func test(e employee) {
// 	fmt.Println(e.getName(), e.getSalary())
// }
// func main() {
// 	test(fullTime{name: "Ahmad Hassan", salary: 5000})
// 	test(contractor{name: "Aladdin", hourlyPay: 587, hoursPerYear: 7884})
// }

// type divideError struct {
// 	dividend float64
// }

// func (de divideError) Error() string {
// 	return fmt.Sprintf("Cannot divide %v by zero", de.dividend)
// }

// func divide(dividend, divisor float64) (float64, error) {
// 	if divisor == 0 {
// 		return 0, divideError{dividend: dividend}
// 	}
// 	return dividend / divisor, nil
// }

// func divide(dividend, divisor float64) (float64, error) { // another way to create custom error
// 	if divisor == 0 {
// 		return 0, errors.New(fmt.Sprintf("Cannot divide %v by zero", dividend))
// 	}
// 	return dividend / divisor, nil
// }

// func main() {
// 	result, err := divide(10, 0)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}
// }

// func main() {
// 	for i := 0; i <= 10; i++ { // for loop example
// 		fmt.Println(i)
// 	}
// }

// func main() { go does not have a while loop just add a condition in the for loop
// 	for true {
// 		fmt.Println("Hello")
// 	 }
// }

// func main() { go does not have a while loop just add a condition in the for loop
// 	counter := 0
// 	for true {
// 		if counter == 5 {
// 			break
// 		}
// 		fmt.Println("Hello")
// 		counter++
// 	}
// }

//  logical AND operator &&
// logical OR operator ||
// https://golangtutorial.dev/tips/go-1.18-build-error-go-linkname-must-refer-to-declared-function-or-variable/
func main() { // for each loop in go, _ is the index and value is the value in the list
	listInGo := []string{"ahmad", "hassan", "mohsen", "mohammad", "ali"}
	for _, value := range listInGo {
		fmt.Println(value)
	}
}
