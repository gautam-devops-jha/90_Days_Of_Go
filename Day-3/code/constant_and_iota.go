package main

import "fmt"

func main() {
	const a = 42

	const b float32 = 3.14

	var i int = a
	var f float64 = a

	fmt.Println(i)
	fmt.Println(f)

	var f32 float32 = b
	fmt.Println(f32)

	// This is not allowed

	// var f64 float64 = b
	// fmt.Println(f64)

	// We can convert it's type and use it if really needed
	var f64 float64 = float64(b)
	fmt.Println(f64)

	const c = iota
	fmt.Println(c)

	const (
		d = 2 * 5
		e             // = 2 * 5
		g = iota      // 2 --> iota related to position
		h             // 3
		i = 10 * iota //  40
	)

	fmt.Println(d, e, g, h, i)

}
