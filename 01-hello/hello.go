package main

import "fmt"

const englishHelloPrefix = "Hello, "
const marathiHelloPrefix = "नमस्कार, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Marathi":
		prefix = marathiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Aniruddha", "Marathi"))
}
