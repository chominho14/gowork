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
}

func printStudent(s Student){
	fmt.Println("성명 : ",s.Name)
	fmt.Println("분반 : ",s.Class)
	fmt.Println("학번 : ",s.Id)
	fmt.Printf("평점 : %.1f\n",s.Score)
}

func defaultStudent(id int) Student{
	var s Student //필드를 사용하기 위해 변수 설정
	s.Id = id
	s.Name = "noname"
	s.Class = "001"
	s.Score = 0.0
	return s
}

type Students []Student

func (s Students) Len() int {return len(s)}//배열을 받아서 Len메소드 받고 int타입으로 하고 리턴 길이
func (s Students) Less(i, j int) bool {return s[i].Score < s[j].Score}//비교하기위해 i,j를 받고 피교하기 위해 bool타입 리턴
func (s Students) Swap(i, j int) { s[i], s[j] = s[j], s[i]}//s[i]에 s[j]를 넣고 s[j]에 s[i]를 넣는다.


func main(){
	s := []Student{
		{"홍길동",1234567,3.91,"A"},
		{Name:"최길동",Id:12341234,Class:"B",Score:3.76,}
		{"박길동", 11112222, 4.01, "C"}
	}
	
	sort.Sort(Students(s))
	fmt.Println(s)
	// var s1 Student = Student{
	// 	"홍길동",
	//  	1234567,
	//   	3.91,
	//    	"A",
	// }

	// var s2 Student = Student{
	// 	Name:"최길동",
	// 	Id:12341234,
	// 	Class:"B",
	// 	Score:3.76,
	//  }
	//  var s3 Student = Student{"박길동", 11112222, 4.01, "C"}


	// printStudent(s[0])
	// printStudent(s[1])
	// printStudent(s[2])
}

// package main

// import "fmt"


// type Student struct{
// 	Name string
// 	Id int
// 	Score float64
// 	Class string
// }

// func printStudent(s Student){
// 	fmt.Println("성명 : ",s.Name)
// 	fmt.Println("분반 : ",s.Class)
// 	fmt.Println("학번 : ",s.Id)
// 	fmt.Printf("평점 : %.1f\n",s.Score)
// }

// // func defaultStudent(id int) Student{
// // 	var s Student //필드를 사용하기 위해 변수 설정
// // 	s.Id = id
// // 	s.Name = "noname"
// // 	s.Class = "001"
// // 	s.Score = 0.0
// // 	return s
// // }

// func defaultStudent(s *Student) {
// 	s.Id = 11112222
// 	s.Name = "noname"
// 	s.Class = "001"
// 	s.Score = 0.0
// }


// func main(){
// 	// var s1 Student = Student{"홍길동", 1234567, 3.91, "A"}
// 	var s1 Student = Student{
// 		"홍길동",
// 	 	1234567,
// 	  	3.91,
// 	   	"A",
// 	}

// 	var s2 Student = Student{
// 		Id:12341234,
// 		Class:"B",
// 	 }

// 	var s3 Student
// 	defaultStudent(&s3)
// 	printStudent(s3)

// 	// s3 := defaultStudent(11112222)
// 	//s3.Class = "C"

// 	printStudent(s3)
// 	printStudent(s1)
// 	printStudent(s2)
// }