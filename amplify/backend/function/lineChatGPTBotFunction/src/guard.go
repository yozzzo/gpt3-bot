package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
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
	lambda.Start(HandleRequest)
}
