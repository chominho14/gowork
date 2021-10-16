package main

import "fmt"

func main(){
	 s1 := []int{1,2,3}
	// s2 := make([]int, len(s1))

	// for i, v := range s1{
	// 	s2[i] = v
	// }

	//s2 := append([]int{}, s1[0], s1[1] , s1[2] )

	s2 := append([]int{}, s1...)

	fmt.Println(s1, s2)
	fmt.Printf("%x %x \n", &s1[0], &s2[0])
}