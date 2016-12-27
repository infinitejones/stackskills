package main

import "fmt"

func main() {

	var large int
	var small int
        var remainder int

	for {
		fmt.Println("Enter a larger number")
		fmt.Scan(&large)
		fmt.Println("Enter a smaller number")
		fmt.Scan(&small)

		if large > small {
			break
		} else {
			fmt.Println("First number must be larger")
		}
	}

	remainder = large % small
	fmt.Println("Remainder is: ", remainder)
}
