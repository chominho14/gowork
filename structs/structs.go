package main

import "fmt"


func main(){
	// 변수 선언과 동시에 구조체 정의
	var s1 struct{
		Name string
		Id int
		Score float64
		Class string
	}

	var s2 struct{
		Name string
		Id int
		Score float64
		Class string
	}

	// var s1 Student = Student{"홍길동", 1234567, 3.91, "A"}
	// var s2 Student
	// s2.Id = 12341234
	// s2.Class = "B"


	fmt.Println("성명 : ",s2.Name)
	fmt.Println("분반 : ",s2.Class)
	fmt.Println("학번 : ",s2.Id)
	fmt.Printf("평점 : %.1f\n",s2.Score)

	fmt.Println("성명 : ",s1.Name)
	fmt.Println("분반 : ",s1.Class)
	fmt.Println("학번 : ",s1.Id)
	fmt.Printf("평점 : %.1f\n",s1.Score)
}


// package main

// import "fmt"


// type Student struct{
// 	Name string
// 	Id int
// 	Score float64
// 	Class string
// }

// func main(){
// 	// var s1 Student = Student{"홍길동", 1234567, 3.91, "A"}
// 	var s1 Student = Student{
// 		"홍길동",
// 	 	1234567,
// 	  	3.91,
// 	   	"A",
// 	}

// 	// var s2 Student
// 	// s2.Id = 12341234
// 	// s2.Class = "B"

// 	var s2 Student = Student{
// 		Id:12341234,
// 		Class:"B",
// 	 }

// 	fmt.Println("성명 : ",s2.Name)
// 	fmt.Println("분반 : ",s2.Class)
// 	fmt.Println("학번 : ",s2.Id)
// 	fmt.Printf("평점 : %.1f\n",s2.Score)

// 	fmt.Println("성명 : ",s1.Name)
// 	fmt.Println("분반 : ",s1.Class)
// 	fmt.Println("학번 : ",s1.Id)
// 	fmt.Printf("평점 : %.1f\n",s1.Score)
// }