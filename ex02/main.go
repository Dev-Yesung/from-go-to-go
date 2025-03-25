package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello, main!")
	fmt.Println(n, err)

	foo, _ := doFoo()
	fmt.Println(foo)

	for i := 0; i < 10; i++ {
		fmt.Println("doFoo", i+1, "times")
		_, _ = doFoo()
	}
}

func doFoo() (int, error) {
	return fmt.Println("I'm in doFoo Function. Not in touFoo")
}
