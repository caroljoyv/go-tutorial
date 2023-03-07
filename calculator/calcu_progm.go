package main

import (
	"fmt"
	"log"

	basic "example.com/operations"
)

func main() {
	fmt.Println("Enter two numbers: ")
	var one, two int
	if _, err := fmt.Scan(&one, &two); err != nil {
		log.Print("Failed to scan one and two due to  ", err)
		return
	}

	fmt.Printf("Sum of %d and %d is %d\n", one, two, basic.Sum(one, two))
	fmt.Printf("Difference of %d and %d is %d\n", one, two, basic.Difference(one, two))
	fmt.Printf("Product of %d and %d is %d\n", one, two, basic.Product(one, two))

}
