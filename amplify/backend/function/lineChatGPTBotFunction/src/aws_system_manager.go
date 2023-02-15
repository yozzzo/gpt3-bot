package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

const (
	REGION  string = "ap-northeast-1"
	PROFILE string = "amplify-gpt3-user"
)

// パラメータストアから設定値取得
func FetchParameterStore(keyName string) (string, error) {

	sess, _ := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(REGION)},
		Profile: PROFILE,
	})
	svc := ssm.New(sess)

	res, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(keyName),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		fmt.Println(err)
		return "Fetch Error", err
	}

	value := *res.Parameter.Value
	return value, nil
}
