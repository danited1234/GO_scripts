package main

import (
	"errors"
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

// func test(e employee) {
// 	fmt.Println(e.getName(), e.getSalary())
// }
// func main() {
// 	test(fullTime{name: "Ahmad Hassan", salary: 5000})
// 	test(contractor{name: "Aladdin", hourlyPay: 587, hoursPerYear: 7884})
// }

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("Cannot divide %v by zero", de.dividend)
}

// func divide(dividend, divisor float64) (float64, error) {
// 	if divisor == 0 {
// 		return 0, divideError{dividend: dividend}
// 	}
// 	return dividend / divisor, nil
// }

func divide(dividend, divisor float64) (float64, error) { // another way to create custom error
	if divisor == 0 {
		return 0, errors.New(fmt.Sprintf("Cannot divide %v by zero", dividend))
	}
	return dividend / divisor, nil
}

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
// func main() { // for each loop in go, _ is the index and value is the value in the list
// 	listInGo := []string{"ahmad", "hassan", "mohsen", "mohammad", "ali"}
// 	fmt.Println(len(listInGo))
// 	for _, value := range listInGo {
// 		fmt.Println(value)
// 	}
// }

func main() {
	for n := 2; n <= 100; n++ {
		if n%2 == 0 {
			fmt.Println(n)
		}
		continue
	}
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a[1:4]) // first index is inclusive and the last index is exclusive
	b := a[:]
	c := []int{78, 9, 9, 2}       // this is a slice of the whole array
	fmt.Println(len(b))           // to get the length of the slice
	b = append(b, 6, 7, 8, 9, 10) // append is used to add elements to the end of the slice
	fmt.Println(b)
	fmt.Println(cap(b)) // to get the capacity of the slice
	fmt.Println(b)
	b = append(b, c...) // appedning slices using another way
	fmt.Println(b)
	i := make([]int, 3, 5) // make is used to create a slice with a specific length and capacity
	fmt.Println(cap(i))
}
func variadicFunction(nums ...string) {
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}

func test() {
	names := []string{"ahmad", "hassan", "mohsen", "mohammad", "ali"}
	variadicFunction(names...)
}

func map_functon() { // map is a collection of key-value pairs just like a dictionary in python
	ages := map[string]int{
		"ahmad":  20,
		"hassan": 21,
		"mohsen": 22,
	}
	fmt.Println(ages["ahmad"])
	names := map[rune]map[string]int{ // this is a nested map an example for a sample json object (rune is a character)
		'a': {"ahmad": 20},
		'b': {"hassan": 21},
		'c': {"mohsen": 22},
	}
	fmt.Println(names['a'])
}
