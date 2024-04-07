package main

import "fmt"

func main() {
	s := "Hello, World"

	p := &s

	fmt.Println(p)
	fmt.Println(*p)

	*p = "Hello Gophers!"

	fmt.Println(s)

	p = new(string)

	fmt.Println(p, *p)

}
