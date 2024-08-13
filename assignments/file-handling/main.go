package main

import (
	"io"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error while opening the file", fileName)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)

}
