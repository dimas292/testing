package main

import (
	"fmt"
)

// belajar pointer sampai bisa
// kapan kita menggunakan pointer
// apa itu interface

type Student struct {
	Id   int
	Name string
	Age  int
}

func (student *Student) setAge(age int) {
	student.Age = age
}

func setStudentName(student *Student, name string) {
	student.Name = name
}

func setStudentID(s *Student, id int){
	s.Id = id
}

func main() {
	var student Student = Student{
		Id:   1,
		Name: "dimas",
		Age:  20,
	}

	fmt.Println("before : ", student.Name)
	
	setStudentName(&student, "jay")
	
	fmt.Println("after : ", student.Name)

	setStudentID(&student, 123)

	fmt.Println("after change id", student.Id)

	student.setAge(19)

	fmt.Println("age after change", student.Age)

	
	fmt.Println("jh")

}
