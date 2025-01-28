package main

import "fmt"

func main() {
	var count int
	fmt.Print("Input your count : ")
	fmt.Scanf("%d", &count)
	for i := 0; i < count; i++ {
		fmt.Println("count : ", i+1)
	}
}
