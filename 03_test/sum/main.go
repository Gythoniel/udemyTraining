package main

import "fmt"

var sum int
var max int

func main() {
	fmt.Print("Maximum: ")
	fmt.Scan(&max)
	for i := 0; i < max; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum is: ", sum)
}
