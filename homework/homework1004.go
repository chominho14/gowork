package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var number1, number2 int
	var temp int
	var arr []float64 = []float64{}
	var sum float64 = 0

	fmt.Print("두 수를 입력해주세요 : ")
	_, err := fmt.Scanf("%d %d", &number1, &number2)

	if err != nil {
		log.Fatal(err)
	}
	if number1 > number2 {
		temp = number1
		number1 = number2
		number2 = temp
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
				arr = append(arr, float64(i))
				break
			}
		}
	}
	fmt.Printf("%.0f ", arr)
	for _, number := range arr {
		sum = sum + number
	}
	fmt.Printf("\n평균 : %.2f", sum/float64(len(arr)))
}
