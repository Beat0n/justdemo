package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Go语言学习园地")
	//Seek设置偏移量
	reader.Seek(-6, os.SEEK_END)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
}
