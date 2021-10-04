package main

import (
	"fmt"
	"math"
	"strings"
) //포맷관련 패키지
//수학관련 패키지
//문자관련 패키지

//진입점
func main() {
	var c rune = '가' //rune은 작은 따옴표
	//var a int16 = 7
	//var a = 7
	a := 7
	var b float64 = 5.3
	a = int(b) //Type Conversion, Casting
	d := false

	fmt.Println(d)
	fmt.Printf("%T", d)
	fmt.Println(a)

	fmt.Println(c)        //유니코드(UTF-8) 값 출력
	fmt.Printf("%c\n", c) //한 글자 출력
	fmt.Printf("%T\n", c) //rune은 결국 int32의 별명이다.

	fmt.Println(math.Round(2.71)) //Ceil 올림처리 함수, Floor 내림처리, Round 반올림
	fmt.Println("Hello Go~")
	fmt.Println(strings.Title("go git github java")) //strings.Title맨 앞 글자들을 대문자로 바꿔준다.

}
