package main

import "fmt"

var a int = 64

type hotdog int

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b hotdog = 100
	fmt.Printf("%T\n", b)

	var i int = int(b)
	fmt.Printf("%T\n", i)
	fmt.Printf("%d\n", i)
}
