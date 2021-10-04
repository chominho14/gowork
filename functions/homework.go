package main

import (
	"fmt"
	"log"
	"os"
)

//두 수를 입력받아 그 사이에 있는 소수들을 출력하는 것
func main() {
	var number1, number2 int
	var k int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanf("%d %d", &number1, &number2)

	if err != nil {
		log.Fatal(err)
	}
	if number1 > number2 {
		k = number1
		number1 = number2
		number2 = k
	}
	if number1 < 2 {
		fmt.Println(number1, "는(은) 소수가 아닙니다~")
		os.Exit(0)
	}

	for i := number1; i <= number2; i++ {
		isPrime := true
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				isPrime = false
			}
			if isPrime {
				fmt.Printf("%d ", i)
				break
			}
		}

	}
}
