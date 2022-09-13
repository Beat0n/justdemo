package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Utf8Index("Go语言学习园地", "学习"))
}

// strings.Index 的 UTF-8 版本
// 即 Utf8Index("Go语言学习园地", "学习") 返回 4，而不是 strings.Index 的 8
func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}
