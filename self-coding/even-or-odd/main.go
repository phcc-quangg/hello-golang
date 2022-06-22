package main

import (
	"fmt"
)

type numbers []int

func main() {
	evenOrOdd()
}

func evenOrOdd() {
	nums := numbers{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range nums {
		if (num % 2) == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
