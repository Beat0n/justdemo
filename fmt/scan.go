package main

import "fmt"

func main() {
	var (
		name string
		age  int
	)
	//Scan/FScan/Sscan 这组函数将连续由空格分隔的值存储为连续的实参（换行符也记为空格）。
	n, _ := fmt.Sscan("polaris 28", &name, &age)
	// 可以将"polaris 28"中的空格换成"\n"试试
	// n, _ := fmt.Sscan("polaris\n28", &name, &age)
	fmt.Println(n, name, age)
}
