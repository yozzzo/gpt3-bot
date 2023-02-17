package main

import (
	"fmt"

	lineBot "github.com/line/line-bot-sdk-go/linebot"
)

// リクエストボディを受け取る構造体
type Response struct {
	RequestBody string `json:"RequestBody"`
}

// リクエストボディから特定のパラメータを受け取る構造体
type Event struct {
	Events []struct {
		Message struct {
			Text string `json:"text"`
		} `json:"message"`
		Source struct {
			UserID string `json:"userId"`
		} `json:"source"`
	}
}

func PostLineMessage(userId string, text string) {
	fmt.Println("line_api呼び出し")
	var LINE_CHANNEL_SECRET, _ = FetchParameterStore("/amplify/d1czxcjz8q1jue/dev/AMPLIFY_lineChatGPTBotFunction_LINE_CHANNEL_SECRET")
	var LINE_CHANNEL_ACCESS_TOKEN, _ = FetchParameterStore("/amplify/d1czxcjz8q1jue/dev/AMPLIFY_lineChatGPTBotFunction_LINE_CHANNEL_ACCESS_TOKEN")

	bot, err := lineBot.New(LINE_CHANNEL_SECRET, LINE_CHANNEL_ACCESS_TOKEN)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := bot.PushMessage(userId, lineBot.NewTextMessage(text)).Do(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("line_api完了")
}

// func handler(request events.APIGatewayProxyRequest) {
// 	// 先ほど定義した構造体
// 	var event Event

// 	// リクエストボディを受け取る
// 	body := request.Body
// 	res := Response{
// 		RequestBody: body,
// 	}

// 	json.Unmarshal([]byte(res.RequestBody), &event)

// 	// 送信元userIDとテキスト内容を変数に代入
// 	userid := fmt.Sprintf("%v", event.Events[0].Source.UserID)
// 	text := fmt.Sprintf("%v", event.Events[0].Message.Text)

// 	// 送信元ユーザにメッセージを返信する関数(次に実装します)
// 	postLineMessage(userid, text)
// }
