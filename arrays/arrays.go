package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5

	//var primes [3]int = [3]int{2, 3, 5}

	primes := [3]int{2, 3, 5}

	//초기화 하지 않은 원소의 제로 값은 해당 원소의 타입의 제로값으로 결정
	//test := [5]bool{true, true, true}
	//fmt.Println(test[3])

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(primes[i])
	// }

	// i := 0
	// for i < len(primes) { //go 언어에서는 len함수가 길이를 말해준다.
	// 	fmt.Println(primes[i])
	// 	i++
	// }

	// for index, prime := range primes {
	// 	fmt.Println(index, prime)
	// }

	// for _, prime := range primes {
	// 	fmt.Println(prime)
	// }

	var sum int = 0
	for _, prime := range primes {
		//fmt.Println(prime)
		sum = sum + prime
	}
	fmt.Println(sum)
	//fmt.Println(float64(sum) / float64(len(primes)))
	fmt.Printf("%.2f\n", float64(sum)/float64(len(primes)))

	//fmt.Println("Arrays")
	//fmt.Println(primes[2])
}
