package main

import (
	"fmt"
)

// map을 사용자 정의로 사용(key는 string, value는 int로 저장.
type vendingMachine map[string]int

// 자판기에 들어있는 값을 채워주는 함수이다.
func(v vendingMachine) Value(value string, price int){
	v[value] = price
}

func main(){
	//타입이 vendingMachine인 v.
	var v vendingMachine
	//make함수를 이용해 vendingMachine타입의 0개의 v를 생성
	v = make(vendingMachine, 0)

	// 두개의 key, value를 입력해준다.
	v.Value("코카콜라",1000)
	v.Value("펩시",900)
	fmt.Println("메뉴 구성")

	//For... range 구문으로 저장되어있는 값 출력
	for key,val := range v {
		fmt.Printf("종류 : %s, 가격 : %d원\n",key, val)
	}
}