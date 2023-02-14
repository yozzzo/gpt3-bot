package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

const (
	REGION string = "ap-northeast-1"
)

func main() {
	// parameterNameをパラメータストアから取得
	parameterName string = "SamplePrameter"
	res, err := fetchParameterStore(parameterName)
	if err != nil {
		fmt.Println(res, err)
		os.Exit(0)
	}

	fmt.Println(res)
}

// パラメータストアから設定値取得
func fetchParameterStore(param string) (string, error) {

	sess := session.Must(session.NewSession())
	svc := ssm.New(
		sess,
		aws.NewConfig().WithRegion(REGION),
	)

	res, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(param),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "Fetch Error", err
	}

	value := *res.Parameter.Value
	return value, nil
}
