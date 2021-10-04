package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int, m int) (bool, error) {
	var k int
	if n > m {
		k = n
		n = m
		m = k
	}
	if n < 2 {
		fmt.Println(n, "는(은) 소수가 아닙니다~")
		return false, fmt.Errorf("%d 는(은) 소수가 아닙니다.", n)
	}

	for i := n; i <= m; i++ {
		prime := true

		for j := 2; j <= i; j++ {
			if i%j == 0 {
				prime = false
			}
			if prime {
				fmt.Printf("%d ", i)
				break
			}
		}

	}

	return
}

//두 수를 입력받아 그 사이에 있는 소수들을 출력하는 것
func main() {
	var number1, number2 int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanf("%d %d", &number1, &number2)

	if err != nil {
		log.Fatal(err)
	}

	p, err := isPrime(number1, number2)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	if p {

	}

}
