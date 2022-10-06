package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type student struct {
	rollnumber int
	name       string
	address    string
	subject    []string
}

func newstudent(rollno int, name string, address string, subject []string) *student {
	s := new(student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.subject = subject
	return s
}

type studentlist struct {
	list []*student
}

func (ls *studentlist) createstudent(rollno int, name string, address string, subject []string) *student {
	st := newstudent(rollno, name, address, subject)
	ls.list = append(ls.list, st)
	return st
}

func (ls *studentlist) display() {
	for i := range ls.list {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("Student Roll-No: %d\n", ls.list[i].rollnumber)
		fmt.Printf("Student Name: %s\n", ls.list[i].name)
		fmt.Printf("Student Address: %s\n\n", ls.list[i].address)
		fmt.Printf("Student subjects: %s\n\n", ls.list[i].subject)
	}
}

func stringcon(rollno int, name string, address string, subject []string) string {
	stroll := strconv.Itoa(rollno)
	stsub := strings.Join(subject, " ")
	block := stroll + name + address + stsub
	return block
}

func main() {
	student := new(studentlist)
	var x = []string{"ENGLISH", "OOP", "DF"}
	var y = []string{"PAK.STD", "INFO", "AI"}
	student.createstudent(1758, "saud", "AAAA", x)
	student.createstudent(2165, "muiz", "BBBB", y)
	student.display()
	block1 := stringcon(1758, "saud", "AAAA", x)
	sum := sha256.Sum256([]byte(block1))
	fmt.Printf("%x\n", sum) //hexadecimal
	
}
