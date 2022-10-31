package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: missing argument file!")
		os.Exit(42)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
