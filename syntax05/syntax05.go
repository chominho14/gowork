package main

import (
	"fmt"
)

// shadowing
func main() {
	//자료타입을 변수명으로 사용
	// var float64 float64 = 9.1
	// var test float64 = 7.9
	// fmt.Println(float64)
	var test1 float64 = 9.4
	var test2 float64 = 4.4
	fmt.Println(test1 + test2)

	// 패키지를 변수명으로 사용
	// var fmt string = "inha"
	// fmt.Println(fmt)
	var univ string = "inha"
	fmt.Println(univ)

	// 함수를 변수명으로 사용
	// var append string = "functions"
	// var f = append([]string{},"함수")
	var f1 string = "function"
	var f2 = append([]string{}, "함수")
	fmt.Println(f1)
	fmt.Println(f2)
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// //입력된 수의 소수 판정 프로그램(0과1처리) v0.3
// func main() {
// 	var number int

// 	fmt.Print("점수 입력 : ")
// 	_, err := fmt.Scanln(&number)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for number < 2 {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 		os.Exit(0)
// 	}

// 	isPrime := true

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}

// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {

// 	}
// }

// package main

// import (
// 	"fmt"
// 	//"math/rand"
// 	//"time"
// 	"bufio"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// //입력된 수의 소수 판정 프로그램 v0.1
// func main() {
// 	fmt.Print("점수 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n')

// 	if err != nil { //err발생하면
// 		log.Fatal(err)
// 	}

// 	in = strings.TrimSpace(in) //불필요한 앞,뒤 제거
// 	number, err := strconv.Atoi(in)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}

// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")

// 	}
// }
