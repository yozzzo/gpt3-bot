package main

import (
	"context"
	"fmt"
)

type MyEvent struct {
	Name string `json:"name"`
}

type APIGatewayProxyResponse struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func HandleRequest(ctx context.Context, name MyEvent) (APIGatewayProxyResponse, error) {
	return APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       `{"greeting": "hello world"}`,
	}, nil
}

func main() {
	// lambda.Start(HandleRequest)
	fmt.Println("起動")
	FetchParameterStore("/amplify/d1czxcjz8q1jue/dev/AMPLIFY_lineChatGPTBotFunction_OPEN_AI_API_KEY")
}
