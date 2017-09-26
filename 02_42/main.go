package main

import "fmt"

func main() {
	for i := 10000000; i < 10000121; i++ {
	fmt.Printf("%d \t %b \t %#x \t %q \n", i, i, i, i)
	}

}
