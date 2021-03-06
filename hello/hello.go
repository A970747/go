package hello

import (
	"fmt"
)

const englishPrefix = "Hello! "
const spanishPrefix = "Hola! "
const frenchPrefix = "Bonjour! "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchPrefix
	case "spanish":
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
