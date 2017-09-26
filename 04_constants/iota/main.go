package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Println("binary\t\t\t\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t\t\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t\t", TB)
	fmt.Printf("%d\n", TB)

}
