package structs

import (
	"fmt"
)

type customer struct {
	firstName string
	lastName string
	age int
}

func (a customer) save() {
	fmt.Println("Added : ", a.firstName)
}

func (a customer) save() {
	fmt.Println("Updated  : ", a.firstName)
}

func Demo2() {
	a := customer{firstname: "Joseph" , lastname: "Smith"
