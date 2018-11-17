package main

import "fmt"

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frechHelloPrefix = "Bonjour, "
const englishPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frechHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}