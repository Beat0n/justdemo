package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("tmp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := io.MultiWriter(os.Stdout, file)
	writer.Write([]byte("MultiWriter!"))
}
