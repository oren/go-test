package main

import "fmt"

// Hello is a Function that returns hello and a name
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
