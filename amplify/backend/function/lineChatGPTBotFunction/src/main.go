package main

import (
	"encoding/json"
	"fmt"

	events "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type APIGatewayProxyResponse struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

// リクエストボディを受け取る構造体
type Responses struct {
	RequestBody string `json:"RequestBody"`
}

// リクエストボディから特定のパラメータを受け取る構造体
type Events struct {
	Events []struct {
		Message struct {
			Text string `json:"text"`
		} `json:"message"`
		Source struct {
			UserID string `json:"userId"`
		} `json:"source"`
	}
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// 先ほど定義した構造体
	var event Events

	// リクエストボディを受け取る
	body := request.Body
	res := Responses{
		RequestBody: body,
	}

	json.Unmarshal([]byte(res.RequestBody), &event)

	// 送信元userIDとテキスト内容を変数に代入
	userId := fmt.Sprintf("%v", event.Events[0].Source.UserID)
	text := fmt.Sprintf("%v", event.Events[0].Message.Text)

	fmt.Println(userId)
	fmt.Println(text)

	// 送信元ユーザにメッセージを返信する関数(次に実装します)
	PostLineMessage(userId, text)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       `{"greeting": "hello world"}`,
	}, nil
}

func main() {
	fmt.Println("起動")
	lambda.Start(HandleRequest)
	// PostOpenAi("美味しいカレーの作り方教えて。返答するときは必ず語尾に「ナン！」とつけてください")
	// PostLineMessage("Ue8d8652379d98fb101ec89c6110331aa", "aaaa")
	// Lineからメッセージを受け取る
	// var messageFromLine = "aaa"
	// OpenAiに流す
	// var aiMessage string = PostOpenAi(messageFromLine)
	// Lineに返却する
	// PostLineBot(aiMessage)
}
