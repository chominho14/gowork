package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {
	if n < 2 {
		fmt.Println(n, "는(은) 소수가 아닙니다.")
		return false, fmt.Errorf("%d 는(은) 소수가 아닙니다.", n)
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil //소수 판정값, 에러 리턴
}

func prime(number int) {
	p, err := isPrime(number)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	if p {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}

func primeRange(a int, b int) {
	if a > b {
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isPrime(i)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		if p {
			fmt.Print(i, " ")
		}
	}
}

//소수 판정 및 구간 소수 판정 프로그램 v1.7
func main() {
	var menu int

	for true {
		fmt.Print("1) 소수판정 2) 구간 소수판정 :")
		_, err := fmt.Scanln(&menu)

		if err != nil {
			log.Fatal(err)
		}

		switch menu {
		case 1:
			var in int
			fmt.Print("정수 1개 입력 : ")
			_, err := fmt.Scanln(&in)
			if err != nil {
				log.Fatal(err)
			}
			prime(in)
		case 2:
			var n1, n2 int
			fmt.Print("정수 2개 입력 : ")
			_, err := fmt.Scanln(&n1, &n2)
			if err != nil {
				log.Fatal(err)
			}
			primeRange(n1, n2)
		default:
			fmt.Print("프로그램을 종료합니다.")
			os.Exit(1)
		}
	}

}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool, error) {
// 	if n < 2 {
// 		fmt.Println(n, "는(은) 소수가 아닙니다.")
// 		return false, fmt.Errorf("%d 는(은) 소수가 아닙니다.", n)
// 	}

// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			return false, nil
// 		}
// 	}
// 	return true, nil //소수 판정값, 에러 리턴
// }

// //구간 소수 판정 프로그램 v1.4 - isPrime함수 안에 변수를 줄이고 return구문 추가, break 제거
// func main() {
// 	var a, b int

// 	fmt.Print("점수 입력 : ")
// 	_, err := fmt.Scanln(&a, &b)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if a > b {
// 		temp := a
// 		a = b
// 		b = temp
// 	}

// 	for i := a; i <= b; i++ {
// 		p, err := isPrime(i)
// 		if err != nil {
// 			log.Fatal(err)
// 			os.Exit(0)
// 		}
// 		if p {
// 			fmt.Print(i, " ")
// 		}
// 	}

// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool, error) {
// 	prime := true
// 	if n < 2 {
// 		fmt.Println(n, "는(은) 소수가 아닙니다.")
// 		return false, fmt.Errorf("%d 는(은) 소수가 아닙니다.", n)
// 	}

// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			prime = false
// 			break
// 		}
// 	}
// 	return prime, nil //소수 판정값, 에러 리턴
// }

// //구간 소수 판정 프로그램 v1.2
// func main() {
// 	var a, b int

// 	fmt.Print("점수 입력 : ")
// 	_, err := fmt.Scanln(&a, &b)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if a > b {
// 		temp := a
// 		a = b
// 		b = temp
// 	}

// 	for i := a; i <= b; i++ {
// 		p, err := isPrime(i)
// 		if err != nil {
// 			log.Fatal(err)
// 			os.Exit(0)
// 		}
// 		if p {
// 			fmt.Print(i, " ")
// 		}
// 	}

// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool, error) {
// 	prime := true
// 	if n < 2 {
// 		fmt.Println(n, "는(은) 소수가 아닙니다.")
// 		return false, fmt.Errorf("%d 는(은) 소수가 아닙니다.", n)
// 	}

// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			prime = false
// 			break
// 		}
// 	}
// 	return prime, nil //소수 판정값, 에러 리턴
// }

// //소수 판정 프로그램 v1.1 : 함수적용, error리턴
// func main() {
// 	var number int

// 	fmt.Print("점수 입력 : ")
// 	_, err := fmt.Scanln(&number)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	p, err := isPrime(number)
// 	if err != nil {
// 		log.Fatal(err)
// 		os.Exit(0)
// 	}

// 	if p {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) bool {
// 	prime := true

// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			prime = false
// 			break
// 		}
// 	}
// 	return prime //true리턴 =소수, false !=소수
// }

// //소수 판정 프로그램 v1.0 : 함수적용
// func main() {
// 	var number int

// 	fmt.Print("점수 입력 : ")
// 	_, err := fmt.Scanln(&number)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if number < 2 {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 		os.Exit(0)
// 	}

// 	if isPrime(number) {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }

//after(multi return)
// package main

// import "fmt"

// func processScore(kor int, eng int, math int) (int, int) {
// 	totalScore := kor + eng + math
// 	average := totalScore / 3

// 	return totalScore, average
// 	//fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", name, totalScore, average)
// }

// func main() {
// 	var t int
// 	var a int

// 	t, a = processScore(100, 90, 93)
// 	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", "홍길동", t, a)

// 	t, a = processScore(94, 93, 91)
// 	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", "홍길순", t, a)

// }

//before
// package main

// import "fmt"

// func main() {
// 	kor := 100
// 	eng := 90
// 	math := 93
// 	name := "홍길동"

// 	fmt.Println(name, "의 총점은", kor+eng+math, "입니다. 평균은", (kor+eng+math)/3.0)

// 	kor = 94
// 	eng = 93
// 	math = 91
// 	name = "홍길순"

// 	fmt.Println(name, "의 총점은", kor+eng+math, "입니다. 평균은", (kor+eng+math)/3.0)

// }
