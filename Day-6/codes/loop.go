package main

import "fmt"

func main() {
	i := 99

	for {
		fmt.Println(i)
		i += 1
		// break ---> used to imidiatly terminate the loop -->so if break statement is used here only 99 is printed to terminal
	}

	// conditional loop
	i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// counter based loop

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
