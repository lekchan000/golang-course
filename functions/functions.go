package main

import "fmt"

func hello() {
	fmt.Println("Hello Suriya")
}

func plus(val1 int8, val2 int8) {
	result := val1 + val2
	fmt.Println("result = ", result)
}

func minus(val3 int8, val4 int8) int8 {
	return val3 + val4
}

func multiply(val5, val6, val7 int16) int16 {
	return val5 * val6 * val7
}

func main() {
	hello()
	plus(1, 4)
	minus := minus(3, 4)
	fmt.Println("result = ", minus)
	mul := multiply(2, 5, 10)
	fmt.Println("result = ", mul)

}
