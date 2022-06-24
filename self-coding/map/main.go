package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["black"] = "#000"

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"blue":  "#00ff00",
	}

	logColor(colors)

	// fmt.Println(colors)
}

func logColor(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
