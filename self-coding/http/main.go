package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("https://google.com/")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Attempt #1 by my own research
	// body, err := io.ReadAll(resp.Body)
	// resp.Body.Close()
	// fmt.Println(string(body))

	// Attempt #2 by teacher with byteslice
	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	//Attempt #3 with io
	lw := logWriter{}
	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("len(bs) = ", len(bs))

	return len(bs), nil
}
