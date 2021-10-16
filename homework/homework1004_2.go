package main

import (
	"fmt"
	"log"
	"os"
)

func Prime(num1, num2 int) ([]float64, error) {
	//리턴할 슬라이스와 처음 입력수가 더 컸을 경우를 생각해서 arr, temp를 만들어 준다.
	var arr []float64 = []float64{}
	var temp int

	//num1과 num2를 입력받는다.
	_, err := fmt.Scanf("%d %d", &num1, &num2)

	//에러일 경우 처리
	if err != nil {
		log.Fatal(err)
	}
	//첫 번째 입력 수가 두 번째 입력 수 보다 클 경우 바꿔주기.
	if num1 > num2 {
		temp = num1
		num1 = num2
		num2 = temp
	}
	//입력 수 중 작은 수가 마이너스 거나 0, 1, 0.23처럼 다른 수 일 경우를 처리.
	if num1 < 2 {
		fmt.Println(num1, "는(은) 소수가 아닙니다~")
		os.Exit(0)
	}
	//입력된 수들 사이에 소수이면 arr슬라이스에 추가해 준다. 
	for i := num1; i <= num2; i++ {
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
	return arr, err
}

func main() {
	var sum float64 = 0
	var n1, n2 int

	fmt.Print("두 수를 입력해주세요 : ")
	arrs, err := Prime(n1, n2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arrs)

	//for...range문을 이용하여 arrs에 있는 값들을 sum에 더해서 넣어준다.
	for _, number := range arrs {
		sum = sum + number
	}
	fmt.Printf("평균 : %.2f", sum/float64(len(arrs)))
}



// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

//version-01 메인함수에서만 만들어 보기.
// func main() {
// 	var number1, number2 int
// 	var temp int
// 	var arr []float64 = []float64{}
// 	var sum float64 = 0

// 	fmt.Print("두 수를 입력해주세요 : ")
// 	_, err := fmt.Scanf("%d %d", &number1, &number2)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if number1 > number2 {
// 		temp = number1
// 		number1 = number2
// 		number2 = temp
// 	}
// 	if number1 < 2 {
// 		fmt.Println(number1, "는(은) 소수가 아닙니다~")
// 		os.Exit(0)
// 	}
// 	for i := number1; i <= number2; i++ {
// 		isPrime := true
// 		for j := 2; j <= i; j++ {
// 			if i%j == 0 {
// 				isPrime = false
// 			}
// 			if isPrime {
// 				arr = append(arr, float64(i))
// 				break
// 			}
// 		}
// 	}
// 	fmt.Printf("%.0f ", arr)
// 	for _, number := range arr {
// 		sum = sum + number
// 	}
// 	fmt.Printf("\n평균 : %.2f", sum/float64(len(arr)))
// }
