package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	//ReadSlice 返回的 []byte 是指向 Reader 中的 buffer ，而不是 copy 一份返回, 因此会被下次IO操作重写
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))

	newreader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line1, _ := newreader.ReadBytes('\n')
	fmt.Printf("the line1 is %s", line1)
	line2, _ := newreader.ReadBytes('\n')
	fmt.Printf("the line2 is %s\n", line2)
	fmt.Printf("the line1 is %s\n", line1)
}
