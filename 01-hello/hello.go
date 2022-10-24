package main

import "fmt"

const englishHelloPrefix = "Hello, "
func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }
    if language == "Marathi" {
        return "नमस्कार, " + name
    }
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Aniruddha", "Marathi"))
}
