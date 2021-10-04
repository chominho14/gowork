package main

import (
	"fmt"
	"math/rand"
	"time"
)

//소수 판정 프로그램 v0.4
func main() {

	//seed 설정
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")

	}
}

// //소수 판정 프로그램 v0.4
// func main() {

// 	//seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	isPrime := true
// 	number := rand.Intn(150) + 2
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 		}
// 	}

// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")

// 	}
// }

// //소수 판정 프로그램 v0.3
// func main() {

// 	//seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	isPrime := true
// 	number := rand.Intn(150) + 2
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 		}
// 	}

// 	if isPrime == true {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")

// 	}
// }

// //소수 판정 프로그램 v0.2
// func main() {

// 	//seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	count := 0
// 	number := rand.Intn(150) + 2 //0과 1 제외 2~151사이의 수
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i < number; i++ { //1과 number일때 loop 안 돔.
// 		if number%i == 0 {
// 			count++
// 		}
// 	}

// 	if count == 0 {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }

// }
// //소수 판정 프로그램 v0.1
// func main() {

// 	//seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	count := 0
// 	number := rand.Intn(150) + 2 //0과 1 제외 2~151사이의 수
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 1; i <= number; i++ {
// 		if number%i == 0 {
// 			count++
// 		}
// 	}

// 	if count == 2 {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")

// 	}
// }
