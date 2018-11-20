package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"strings"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	text := request.QueryStringParameters["text"]
	data, err := ioutil.ReadFile("template.txt")

	newText := strings.Replace(string(data), ":nhin_cai_gi:", text, -1)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: newText,
	}, nil
}

func main() {
	lambda.Start(handler)
}
