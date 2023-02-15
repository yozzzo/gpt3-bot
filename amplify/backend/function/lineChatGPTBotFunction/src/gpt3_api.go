package main

// import (
// 	"context"
// 	"fmt"

// 	gogpt "github.com/sashabaranov/go-gpt3"
// )

// const (
// 	GPT3_COMPLETIONS_ENDPOINT string = "https://api.openai.com/v1/completions"

// 	// # Model name
// 	GPT3_MODEL string = "text-davinci-003"

// 	// # API request headers
// 	// string HEADERS = {
// 	// 	"Content-Type": "application/json",
// 	// 	"Authorization": f"Bearer {const.OPEN_AI_API_KEY}"
// 	// }

// 	// # Maximum number of tokens to generate
// 	MAX_TOKENS int = 1000

// 	// # Controls the randomness of the generated text
// 	TEMPERATURE float32 = 0.5

// 	// # Number of completions to generate
// 	GENERATE_COMPLETIONS_COUNT int = 1

// 	// # Specifies the token at which to stop generating completions
// 	STOP bool = false
// )

// func main() {
// 	c := gogpt.NewClient("your token")
// 	ctx := context.Background()

// 	req := gogpt.CompletionRequest{
// 		Model:     gogpt.GPT3Ada,
// 		MaxTokens: 5,
// 		Prompt:    "Lorem ipsum",
// 	}
// 	resp, err := c.CreateCompletion(ctx, req)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(resp.Choices[0].Text)
// }
