package main

import (
	"fmt"
	"time"
)

type student struct{
	id int 
	name string
	subject string
	seconds int
}

func (s student) study(){
	fmt.Printf("%s은 %s 과목을 공부합니다.",s.name,s.subject)
}

func (s student) countDown(){
	for s.seconds > 0 {
		fmt.Println(s.seconds)
		time.Sleep(time.Second)
		s.seconds--
	}
}

func funcStudy(fs student){
	fmt.Printf("%s은 %s 과목을 공부합니다.",fs.name,fs.subject)
}

func main(){
	s1 := student{1000, "홍길동", "Go",10}
	//fmt.Println(s1)
	//funcStudy(s1)
	s1.countDown()
	s1.study()
}

// package main

// import "fmt"

// type Liters float64
// type Gallons float64
// type Milliliters float64

// func (g Gallons)ToLiters() Liters {
// 	return Liters(g * 3.785)
// }

// func (g Gallons)GallonsToMilliliters() Liters {
// 	return Liters(g * 3785.41)
// }


// func (l Liters)ToGallons() Gallons {
// 	return Gallons(l * 0.264)
// }

// func (ml Milliliters)ToGallons() Gallons{
// 	return Gallons(ml * 0.000264)
// }

// func main(){
// 	//coke := Liters(2)
// 	var coke Liters = 2
// 	fanta := Milliliters(750)
// 	fmt.Printf("%.2f리터는 %0.3f갤론과 같습니다.\n",coke,coke.ToGallons())
// 	fmt.Printf("%.2f밀리리터는 %0.3f갤론과 같습니다.\n",fanta,fanta.ToGallons())

// }



// package main

// import "fmt"

// type Number int

// func (n *Number) Square()  {
// 	*n= *n * *n
// 	//return n
// }

// func pointerSquare(n *int){
// 	*n = *n * *n
// }

// func main(){
// 	n1 := Number(5)
// 	n2 := 4
// 	//fmt.Println(n1.Square())

// 	pointerSquare(&n2)
// 	fmt.Println(n2)
	
// 	fmt.Println(n1)
// 	n1.Square()
// 	fmt.Println(n1)

// }


// package main

// import "fmt"

// type Liters float64
// type Gallons float64
// type Milliliters float64

// func GallonsToLiters(g Gallons) Liters {
// 	return Liters(g * 3.785)
// }

// func GallonsToMilliliters(g Gallons) Liters {
// 	return Liters(g * 3785.41)
// }


// func LitersToGallons(l Liters) Gallons {
// 	return Gallons(l * 0.264)
// }

// func MillilitersToGallons(ml Milliliters) Gallons{
// 	return Gallons(ml * 0.000264)
// }

// // func ToLiters(g Gallons) Liters {
// // 	return Liters(g * 3.785)
// // }
// func main(){
// 	var carFuel Liters = 10
// 	var busFuel Gallons = 60

// 	carFuel = carFuel + ToLiters(Gallons(10.0))
// 	busFuel = busFuel + ToGallons(Liters(10.0))


// 	fmt.Println(carFuel, busFuel)



// }

// package main

// import "fmt"

// type Liters float64
// type Gallons float64
// type Name string

// func main(){
// 	// var carfuel Liters = 10
// 	// var busfuel Gallons = 60
// 	carfuel := Liters(10)
// 	busfuel := Gallons(60)

// 	// carfuel = Liters(Gallons(10)* 3.785)
// 	// busfuel = Gallons(Liters(60)* 0.264)
// 	// //carfuel = Liters(15.9)
// 	// //carfuel = Gallons(15.9) //불가능
// 	// //busfuel = Gallons(50.5)
// 	// fmt.Println(carfuel, busfuel)
// 	// fmt.Println("defined type")

// 	// // fmt.Println(Gallons(1.2) + Gallons(5.5))
// 	// // fmt.Println(Gallons(1.2) - Gallons(5.5))
// 	// // fmt.Println(Gallons(1.2) / Gallons(5.5))
// 	// // fmt.Println(Gallons(1.2) * Gallons(5.5))
// 	// // fmt.Println(Gallons(1.2) == Gallons(5.5))
// 	// // fmt.Println(Gallons(1.2) != Gallons(5.5))
// 	// // fmt.Println(Gallons(1.2) > Gallons(5.5))
// 	// // fmt.Println(Gallons(1.2) < Gallons(5.5))

// 	// fmt.Println(Name("홍길동") + Name("홍길순"))
// 	// //fmt.Println(Name("홍길동") - Name("홍길순"))
// 	// fmt.Println(Name("홍길동") > Name("홍길순"))
// 	// fmt.Println(Name("홍길동") < Name("홍길순"))
// 	// fmt.Println(Name("홍길동") == Name("홍길순"))
// 	// fmt.Println(Name("홍길동") != Name("홍길순"))

// }