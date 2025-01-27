package main

import "fmt"

var numInt, numInt2 = 1000, 2000
var string1 = "hello"

func main() {
	numfloat := 2.4
	string2 := "Good day"
	fmt.Println(numInt, numInt2, numfloat, string2)

	fmt.Println(numInt + numInt2)
	fmt.Println(numfloat + float64(numInt))
	fmt.Println("It is a " + string2)
	fmt.Println("My money = ", numInt2)
}
