package main

import (
	"fmt"
	"gowork/packages/mymath"
)

func main() {
	mymath.Test()
	fmt.Println(mymath.MyPower(2, 10))
	fmt.Println("hi")
	fmt.Println(mymath.MyAbs(-99))
}
