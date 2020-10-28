package main

import "fmt"

// Hello world string
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("yesid"))
}
