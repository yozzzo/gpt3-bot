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
	// PostOpenAi("美味しいカレーの作り方教えて。返答するときは必ず語尾に「ナン！」とつけてください")
	PostLineMessage("Ue8d8652379d98fb101ec89c6110331aa", "aaaa")
	// Lineからメッセージを受け取る
	// var messageFromLine = "aaa"
	// OpenAiに流す
	// var aiMessage string = PostOpenAi(messageFromLine)
	// Lineに返却する
	// PostLineBot(aiMessage)
}
