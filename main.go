package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strings"
)

var heart = `:x::x::x::x::x::x::x::x::x::x::x:
:x::x::x:            :x:           :x::x::x:
:x::x:                                        :x::x:
:x::x::x:                             :x::x::x:
:x::x::x::x:                  :x::x::x::x:
:x::x::x::x::x:      :x::x::x::x::x:
:x::x::x::x::x::x::x::x::x::x::x:`

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	emoji := request.QueryStringParameters["text"]
	newHeart := strings.Replace(heart, ":x:", emoji, -1)

	responseBody, _ := json.Marshal(map[string]string {
		"response_type": "in_channel",
		"text": newHeart,
	})

	return events.APIGatewayProxyResponse{
		Headers: map[string]string {
			"content-type": "application/json",
		},
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}

func main() {
	lambda.Start(handler)
}
