package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const hungarianHelloPrefix = "Szevasz, "
const spanish = "Spanish"
const french = "French"
const hungarian = "Hungarian"
const defaultName = "World!"

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}
	prefix := englishHelloPrefix
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case hungarian:
		prefix = hungarianHelloPrefix
	}
	return prefix + name
}
