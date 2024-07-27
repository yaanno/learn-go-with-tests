package helloworld

const (
	spanish   = "Spanish"
	french    = "French"
	hungarian = "Hungarian"

	hungarianHelloPrefix = "Szevasz, "
	frenchHelloPrefix    = "Bonjour, "
	spanishHelloPrefix   = "Hola, "
	englishHelloPrefix   = "Hello, "
	defaultName          = "World!"
)

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case hungarian:
		prefix = hungarianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
