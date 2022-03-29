package main

import "fmt"

func main() {
	for i := 0; i <= 1000; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
		fmt.Println(i)
	}
	fmt.Println("Done!")
}
