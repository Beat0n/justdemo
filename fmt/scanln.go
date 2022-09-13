package main

import "fmt"

func main() {
	var (
		name string
		age  int
	)
	n, _ := fmt.Sscanln("polaris 28", &name, &age)
	fmt.Println(n, name, age)
}
