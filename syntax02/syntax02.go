package main

import (
	"fmt"
	//"reflect"
	"strings"
)

//진입점
func main() {

	//zero value
	// var b float64
	// var c rune = '가' //rune은 작은 따옴표

	// a := 7

	// fmt.Println(b)

	// fmt.Println(reflect.TypeOf(c))
	// fmt.Println(reflect.TypeOf(a))

	texts := "G@ M@new~~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)
}
