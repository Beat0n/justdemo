package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := io.TeeReader(strings.NewReader("Feel so happy"), os.Stdout)
	buf := make([]byte, 20)
	reader.Read(buf)
	fmt.Println("\n" + string(buf))
}
