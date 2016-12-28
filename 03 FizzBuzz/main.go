package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {

		switch {
			case i % 3 == 0:
				if i % 5 == 0 {
					fmt.Println(i, "FizzBuzz")
				} else {
					fmt.Println(i, "Fizz")
				}
			case i % 5 == 0:
				fmt.Println(i, "Buzz")
			default:
				fmt.Println(i)
		}
	}
}
