package main

import (
	"fmt"
	"strconv"
)

const (
	PRODUCT = "Mobile"
)

func main() {
	fmt.Print("test VAR \n")

	var intOne int = 1
	var intTwo = 2
	intQuick := 3

	fmt.Printf("int Nomal set one %d two %d \nquick set %d \n ", intOne, intTwo, intQuick)

	count(intOne, intTwo)
	intQuick = 7
	typeChange(intQuick)
	///Constants
	const st string = "show static"
	fmt.Println("For  " + st)
	//
	fmt.Println("For  " + PRODUCT)

}

func count(c1 int, c2 int) {

	fmt.Print("count test \n")
	fmt.Println(c1 + c2)
	fmt.Println(c1 - c2)
	fmt.Println(c1 * c2)
	fmt.Println(c1 / c2)
	fmt.Println(c1 % c2)
}

func typeChange(c1 int) {
	testType := 6.8
	showfloat64 := testType / float64(c1)
	fmt.Printf("%f \n", showfloat64)
	var s string = strconv.FormatFloat(showfloat64, 'E', -1, 32)
	fmt.Println("add float64 " + s)

	testString := fmt.Sprintf("%v", showfloat64)
	fmt.Println(testString)

}
