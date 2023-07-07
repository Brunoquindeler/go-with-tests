package hello

const (
	englishHelloPrefix             = "Hello, "
	spanishHelloPrefix             = "Hola, "
	frenchHelloPrefix              = "Bonjour, "
	portugueseBrazilianHelloPrefix = "Olá, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	case "pt-br":
		prefix = portugueseBrazilianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
