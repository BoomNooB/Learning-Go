package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		found := false

		if i%3 == 0 {
			fmt.Print("Fizz")
			found = true
		}

		if i%5 == 0 {
			fmt.Print("Buzz")
			found = true
		}

		if !found {
			fmt.Print(i)
		}

		fmt.Println()
	}
}
