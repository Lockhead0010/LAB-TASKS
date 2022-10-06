// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

type student struct {
	rollnumber int
	name       string
	address    string
}

func newstudent(rollno int, name string, address string) *student {
	s := new(student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type studentlist struct {
	list []*student
}

func (ls *studentlist) createstudent(rollno int, name string, address string) *student {
	st := newstudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

func (ls *studentlist) display() {
	for i := range ls.list {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("Student Roll-No: %d\n", ls.list[i].rollnumber)
		fmt.Printf("Student Name: %s\n", ls.list[i].name)
		fmt.Printf("Student Address: %s\n\n", ls.list[i].address)
	}
}

func main() {
	student := new(studentlist)
	student.createstudent(1758, "saud", "AAAA")
	student.createstudent(2165, "muiz", "BBBB")
	student.display()
}
