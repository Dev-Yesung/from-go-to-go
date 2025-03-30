package main

import "fmt"

/*
use var to declare three variables
the variables should have package level scope
*/
var x int
var y string
var z bool

// create your own type. type be an int
type myInt int

var xx myInt
var yy int

func one() {
	// using the short declaration operator
	x := 42
	y := "James Bond"
	z := true

	// print the values
	// using a single print statement
	fmt.Printf("%d %s %t\n", x, y, z)
	fmt.Println(x, y, z)

	// using multiple print statement
	fmt.Printf("%d\n", x)
	fmt.Printf("%s\n", y)
	fmt.Printf("%t\n", z)
}
