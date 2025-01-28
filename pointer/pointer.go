package main

import "fmt"

func zeropointer(ipointer *int) {
	*ipointer = 0
}

func main() {
	i := 1
	fmt.Println("i =", i)
	zeropointer(&i)
	fmt.Println("i =", i)

}
