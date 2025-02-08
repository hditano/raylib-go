package main

import "fmt"

func myPointer(b *int) {
	*b = *b + 1
}

func main() {
	a := 10
	var xPtr *int = &a

	fmt.Println(a)
	fmt.Println(xPtr)
	myPointer(xPtr)
	fmt.Println(a)
}
