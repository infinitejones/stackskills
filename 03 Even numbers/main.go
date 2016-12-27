package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		if i % 3 == 0 {
			if i % 5 == 0 {
				fmt.Println(i, ": FizzBuzz")
			} else {
				fmt.Println(i, ": Fizz")
			}
		}

		if i % 5 == 0 {
				fmt.Println(i,": Buzz")
			} else {
				fmt.Println(i)
			}
	}
}

