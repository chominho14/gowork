package main

import (
	"fmt"
	"gowork/chap10/calender"
	"log"
)




func main(){
	//date := calender.Date{Year: 2021, Month: 11, Day: 15}
	date := calender.Date{}
	fmt.Println(date)
	err := date.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Println(date.GetDay())
	
}