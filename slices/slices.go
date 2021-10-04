package main

import "fmt"

func fa(a2 [3]int) {
	a2[1] = 7
}

func fs(s2 []int) {
	s2[1] = 7
}

func main() {
	a := [...]int{2, 3, 5} //배열
	s := []int{2, 3, 5}    //slice
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", s)

	a1 := [3]int{2, 3, 5}
	s1 := []int{2, 3, 5}

	fa(a1)
	fs(s1)
	// a[1] = 7
	// s[1] = 7
	fmt.Println(a1, s1)

}

// package main

// import "fmt"

// func main() {

// 	//slices
// 	// var s []int
// 	// s = make([]int, 5)

// 	s := make([]int, 5)

// 	//var s []int = []int{0, 0, 0, 0, 0}

// 	//s := []int{0, 0, 0, 0, 0}

// 	if len(s) == 0 {
// 		fmt.Println("슬라이스가 비어있습니다.")
// 		fmt.Println(s)
// 	}
// 	s[4] = 11
// 	s = append(s, -9, 10, 2)
// 	//fmt.Println(s[0])
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(s[5])

// 	//const로 정하면 값을 바꿀 수 없다.
// 	// const an int = 5
// 	// //var an int =5
// 	// var a [an]int = [5]int{2, 3, 1, 5}
// 	// a[4] = 77
// 	//fmt.Println(a)
// }
