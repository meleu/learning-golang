// all comments represent something I learned

// the 'main' package becomes an executable (not sure if it's accurate)
package main

// needed to use Println
import "fmt"

// grouping constants
const (
	spanish    = "Spanish"
	french     = "French"
	portuguese = "Portuguese"

	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Olá, "
)

// public function starts with an Uppercase letter
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// - private function starts with Lowercase letter
// - 'prefix' is a named return value
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
