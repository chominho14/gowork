package main

import (
	"bufio"
	"fmt"
	"log" //Fatal함수를 사용하기 위해
	"os"
	"strconv" //TrimSpace함수 사용 위해
	"strings" //int로 parse하기 위해
)

func main() {

	fmt.Print("단 입력 : ")
	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n')

	if err != nil { //err발생하면
		log.Fatal(err)
	}

	in = strings.TrimSpace(in) //불필요한 앞,뒤 제거
	dan, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}

	i := 1
	for i < 10 {
		fmt.Printf("%d * %d = %d\n", dan, i, (dan * i))
		i++
	}

}

//-----------------------------------------------

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log" //Fatal함수를 사용하기 위해
// 	"os"
// 	"strconv" //TrimSpace함수 사용 위해
// 	"strings" //int로 parse하기 위해
// )

// func main() {

// 	fmt.Print("단 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n')

// 	if err != nil { //err발생하면
// 		log.Fatal(err)
// 	}

// 	in = strings.TrimSpace(in) //불필요한 앞,뒤 제거
// 	//dan, err := strconv.ParseInt(in, 10, 32)
// 	dan, err := strconv.Atoi(in)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for i := 1; i < 10; i++ {
// 		//fmt.Println(dan, " * ", i, " = ", (dan * i))
// 		fmt.Printf("%d * %d = %d\n", dan, i, (dan * i))
// 	}

// }

//----------------------------------------

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {

// 	fmt.Print("숫자 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, _ := rd.ReadString('\n')
// 	fmt.Println(in)
// }
