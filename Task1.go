// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type person struct {
	name       string
	age        int
	occupation string
}

func main() {
	var per1 person

	per1.name = "Saud"
	per1.age = 22
	per1.occupation = "Student"

	fmt.Println("Name: ", per1.name)
	fmt.Println("Age: ", per1.age)
	fmt.Println("Job: ", per1.occupation)

}
