package main

import (
	"fmt"
	"strings"
)

const spanish = "spanish"
const french = "french"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello world string
func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(lang) + name
}

// getPrefix acording to languaje
func getPrefix(lang string) (prefix string) {
	switch strings.ToLower(lang) {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("yesid", ""))
}
