package main

import "fmt"

var scores int

func main() {
	fmt.Print("Input your scores : ")
	fmt.Scanf("%d", &scores)
	fmt.Println("scores = ", scores)

	if scores >= 90 && scores <= 100 {
		fmt.Println("Your Grade is A")
	} else if scores >= 80 && scores <= 90 {
		fmt.Println("Your Grade is B")
	} else if scores >= 70 && scores <= 80 {
		fmt.Println("Your Grade is C")
	} else if scores >= 60 && scores <= 70 {
		fmt.Println("Your Grade is D")
	} else {
		fmt.Println("Your Grade is F")
	}
}
