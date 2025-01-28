package main

import "fmt"

//make maps
var product = make(map[string]int32)

func main() {
	fmt.Println("product = ", product)
	//Add maps
	product["iMac"] = 30000
	product["iPhone"] = 20000
	product["iPad"] = 10000
	product["iPod"] = 5000
	fmt.Println("iMac", product["iMac"])
	fmt.Println("iPhone", product["iPhone"])
	fmt.Println("iPad", product["iPad"])
	fmt.Println("iPod", product["iPod"])
	fmt.Println("product all = ", product)

	//delete
	delete(product, "iMac")
	delete(product, "iPhone")
	fmt.Println("product del = ", product)

	//modify
	product["iPad"] = 25000
	product["iPod"] = 4000
	fmt.Println("product mod = ", product)

	//useshortform
	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println("Show : ", courseName["101"])
}
