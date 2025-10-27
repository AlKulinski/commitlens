package domain

type ModelProvider string

const (
	OpenAIModelProvider ModelProvider = "openai"
	GroqModelProvider   ModelProvider = "groq"
)

func ParseModelProvider(provider string) ModelProvider {
	switch provider {
	case "openai":
		return OpenAIModelProvider
	case "groq":
		return GroqModelProvider
	default:
		return OpenAIModelProvider
	}
}
