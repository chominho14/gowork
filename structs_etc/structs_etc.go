package main

import (
	"fmt",
	"sort"
)

type Student struct{
	Name string
	Id int
	Score float64
	Class string
	Level int
}

type ScholarshipStudent struct{
	//StudentInfo Student
	Student
	Level int
	Scholarship int
}


func main(){
	s1  := ScholarshipStudent{
		Student{"홍길동",1234567,3.91,"A" ,4},
		1,
		1000000,
	}
	// fmt.Println(s1)
	// fmt.Println(s1.Scholarship)
	// fmt.Println(s1.StudentInfo.Id)

	fmt.Print("장학금 수여 학생명 : %s, 장학금액 : %d\n",s1.Name, s1.Scholarship)
	fmt.Println(s1.Student.Level)
	// s := []Student{
	// 	{"홍길동",1234567,3.91,"A"},
	// 	{Name:"최길동",Id:12341234,Class:"B",Score:3.76,}
	// 	{"박길동", 11112222, 4.01, "C"}
	// }
	
}
