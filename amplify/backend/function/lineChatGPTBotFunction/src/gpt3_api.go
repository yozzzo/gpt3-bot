package main

import (
	"context"
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
)

const (
	GPT3_COMPLETIONS_ENDPOINT string = "https://api.openai.com/v1/completions"
	// # Model name
	// GPT3_MODEL string = "text-davinci-003"
	GPT3_MODEL string = "GPT3Curie"
	// # API request headers
	// string HEADERS = {
	// 	"Content-Type": "application/json",
	// 	"Authorization": f"Bearer {const.OPEN_AI_API_KEY}"
	// }
	// # Maximum number of tokens to generate
	MAX_TOKENS int = 1000
	// # Controls the randomness of the generated text
	TEMPERATURE float32 = 0.5
	// # Number of completions to generate
	GENERATE_COMPLETIONS_COUNT int = 1
	// # Specifies the token at which to stop generating completions
	STOP bool = false
)

func PostOpenAi(prompt string) string {
	fmt.Println("gpt3_api")
	// var BASE_SECRET_PATH = os.Getenv("BASE_SECRET_PATH")
	// var TOKEN, _ = FetchParameterStore(BASE_SECRET_PATH + "OPEN_AI_API_KEY")
	var TOKEN, _ = FetchParameterStore("/amplify/d1czxcjz8q1jue/dev/AMPLIFY_lineChatGPTBotFunction_OPEN_AI_API_KEY")
	c := gogpt.NewClient(TOKEN)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 100,
		Prompt:    prompt,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "Error"
	}
	fmt.Println(resp.Choices[0].Text)
	return resp.Choices[0].Text
}
