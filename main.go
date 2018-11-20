package main

import (
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
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       newHeart,
	}, nil
}

func main() {
	lambda.Start(handler)
}
