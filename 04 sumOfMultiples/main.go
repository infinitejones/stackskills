package main

import "fmt"

func main() {

	var sum int

	for i := 1; i < 1000; i++ {

			if i % 3 == 0 || i % 5 == 0 {
				fmt.Println("i: ",i)
				fmt.Println("sum: ", sum)
				sum += i}
				}
	fmt.Println(sum)
}

