package main

import (
	"bytes"
	"os"
)

func main() {
	reader := bytes.NewReader([]byte("Go语言学习园地"))
	reader.WriteTo(os.Stdout)
}
