package main

import "os"

func main() {
	filenames := []string{"tmp1", "tmp2", "tmp3"}
	data := []byte("I am who")
	Tee(filenames, data)
}

func Tee(filenames []string, data []byte) {
	//文件存在则添加内容，不存在则创建
	for _, filename := range filenames {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			file, err := os.Create(filename)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			file.Write(data)
		} else {
			file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				panic(err)
			}
			file.Write(data)
		}
	}
}
