package main

import (
	"fmt"
)

func main() {
	fmt.Println("================ Ninja level 1-1 ================")
	one()

	fmt.Println("================ Ninja level 1-2 ================")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("================ Ninja level 1-3 ================")
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%v\t%v\t%v\n", x, y, z)
	fmt.Println(s)

	fmt.Println("================ Ninja level 1-4 ================")
	fmt.Println(xx)
	fmt.Printf("%T\n", xx)
	xx = 42
	fmt.Println(xx)

	fmt.Println("================ Ninja level 1-5 ================")
	yy = int(xx)
	fmt.Println(yy)
	fmt.Printf("%T", yy)
}
