package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string { //public function starting with uppercase letter

	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

//using if statements

// if language == spanish{
// 	return SpanishHelloPrefix + name
// }
// if language == french{
// 	return FrenchHelloPrefix + name
// }

// return EnglishHelloPrefix + name

//using switch statement

// prefix := englishHelloPrefix

// switch language {
// case spanish:
// 	prefix = spanishHelloPrefix
// case french:
// 	prefix = frenchHelloPrefix
// }
// return prefix + name

func greetingPrefix(language string) (prefix string) { //private function starting with lowercase letter
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix //when the language is neither Spanish nor French
	}
	return
}

func main() {
	// fmt.println(Hello())
	fmt.Println(Hello("Cristiano", "English"))
}
