package main

import "fmt"

func Namevar() string {
	a := "code"
	b := "wa.rs"
	name := a + b
	return name
}

func main() {
	fmt.Println(Namevar())
}
