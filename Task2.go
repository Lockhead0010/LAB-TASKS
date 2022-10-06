// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	position string
}

type Company struct {
	companyName string
	employees   []Employee
}

func main() {

	e1 := Employee{"Saud", 100, "Student"}
	e2 := Employee{"Hassan", 200, "Student"}
	e3 := Employee{"Muiz", 300, "Student"}

	emplys := []Employee{e1, e2, e3}
	comp := Company{"DIGILLAMA", emplys}
	fmt.Printf("COMPANY DETAILS: %v", comp)
}
